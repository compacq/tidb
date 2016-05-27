package evaluator

import (
	"github.com/juju/errors"
	"github.com/pingcap/tidb/context"
	"github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/parser/opcode"
	"github.com/pingcap/tidb/util/types"
)

func builtinAndAnd(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	leftDatum := args[0]
	rightDatum := args[1]
	if leftDatum.Kind() != types.KindNull {
		var x int64
		x, err = leftDatum.ToBool()
		if err != nil {
			err = errors.Trace(err)
			return
		} else if x == 0 {
			// false && any other types is false
			d.SetInt64(x)
			return
		}
	}
	if rightDatum.Kind() != types.KindNull {
		var y int64
		y, err = rightDatum.ToBool()
		if err != nil {
			err = errors.Trace(err)
			return
		} else if y == 0 {
			d.SetInt64(y)
			return
		}
	}
	if leftDatum.Kind() == types.KindNull || rightDatum.Kind() == types.KindNull {
		d.SetNull()
		return
	}
	d.SetInt64(int64(1))
	return
}

func builtinOrOr(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	leftDatum := args[0]
	rightDatum := args[1]
	if leftDatum.Kind() != types.KindNull {
		var x int64
		x, err = leftDatum.ToBool()
		if err != nil {
			err = errors.Trace(err)
			return
		} else if x == 1 {
			// false && any other types is false
			d.SetInt64(x)
			return
		}
	}
	if rightDatum.Kind() != types.KindNull {
		var y int64
		y, err = rightDatum.ToBool()
		if err != nil {
			err = errors.Trace(err)
			return
		} else if y == 1 {
			d.SetInt64(y)
			return
		}
	}
	if leftDatum.Kind() == types.KindNull || rightDatum.Kind() == types.KindNull {
		d.SetNull()
		return
	}
	d.SetInt64(int64(0))
	return
}

func builtinLogicXor(args []types.Datum, _ context.Context) (d types.Datum, err error) {
	leftDatum := args[0]
	righDatum := args[1]
	if leftDatum.Kind() == types.KindNull || righDatum.Kind() == types.KindNull {
		d.SetNull()
		return
	}
	x, err := leftDatum.ToBool()
	if err != nil {
		err = errors.Trace(err)
		return
	}

	y, err := righDatum.ToBool()
	if err != nil {
		err = errors.Trace(err)
		return
	}
	if x == y {
		d.SetInt64(zeroI64)
	} else {
		d.SetInt64(oneI64)
	}
	return
}

func compareFuncFactory(op opcode.Op) BuiltinFunc {
	return func(args []types.Datum, _ context.Context) (d types.Datum, err error) {
		a, b := types.CoerceDatum(args[0], args[1])
		if a.Kind() == types.KindNull || b.Kind() == types.KindNull {
			// for <=>, if a and b are both nil, return true.
			// if a or b is nil, return false.
			if op == opcode.NullEQ {
				if a.Kind() == types.KindNull && b.Kind() == types.KindNull {
					d.SetInt64(oneI64)
				} else {
					d.SetInt64(zeroI64)
				}
			} else {
				d.SetNull()
			}
			return
		}

		n, err := a.CompareDatum(b)

		if err != nil {
			err = errors.Trace(err)
			return
		}
		var result bool
		switch op {
		case opcode.LT:
			result = n < 0
		case opcode.LE:
			result = n <= 0
		case opcode.EQ, opcode.NullEQ:
			result = n == 0
		case opcode.GT:
			result = n > 0
		case opcode.GE:
			result = n >= 0
		case opcode.NE:
			result = n != 0
		default:
			return d, ErrInvalidOperation.Gen("invalid op %v in comparision operation", op)
		}
		if result {
			d.SetInt64(oneI64)
		} else {
			d.SetInt64(zeroI64)
		}
		return
	}
}

func bitOpFactory(op opcode.Op) BuiltinFunc {
	return func(args []types.Datum, _ context.Context) (d types.Datum, err error) {
		a, b := types.CoerceDatum(args[0], args[1])

		if a.Kind() == types.KindNull || b.Kind() == types.KindNull {
			d.SetNull()
			return
		}

		x, err := a.ToInt64()
		if err != nil {
			err = errors.Trace(err)
			return
		}

		y, err := b.ToInt64()
		if err != nil {
			err = errors.Trace(err)
			return
		}

		// use a int64 for bit operator, return uint64
		switch op {
		case opcode.And:
			d.SetUint64(uint64(x & y))
		case opcode.Or:
			d.SetUint64(uint64(x | y))
		case opcode.Xor:
			d.SetUint64(uint64(x ^ y))
		case opcode.RightShift:
			d.SetUint64(uint64(x) >> uint64(y))
		case opcode.LeftShift:
			d.SetUint64(uint64(x) << uint64(y))
		default:
			err = ErrInvalidOperation.Gen("invalid op %v in bit operation", op)
			return
		}
		return
	}
}

func arithmeticFuncFactory(op opcode.Op) BuiltinFunc {
	return func(args []types.Datum, _ context.Context) (d types.Datum, err error) {
		a, err := coerceArithmetic(args[0])
		if err != nil {
			err = errors.Trace(err)
			return
		}

		b, err := coerceArithmetic(args[1])
		if err != nil {
			err = errors.Trace(err)
			return
		}

		a, b = types.CoerceDatum(a, b)
		if a.Kind() == types.KindNull || b.Kind() == types.KindNull {
			d.SetNull()
			return
		}

		switch op {
		case opcode.Plus:
			return computePlus(a, b)
		case opcode.Minus:
			return computeMinus(a, b)
		case opcode.Mul:
			return computeMul(a, b)
		case opcode.Div:
			return computeDiv(a, b)
		case opcode.Mod:
			return computeMod(a, b)
		case opcode.IntDiv:
			return computeIntDiv(a, b)
		default:
			err = ErrInvalidOperation.Gen("invalid op %v in arithmetic operation", op)
			return
		}
	}
}

func unaryOpFactory(op opcode.Op) BuiltinFunc {
	return func(args []types.Datum, _ context.Context) (d types.Datum, err error) {
		defer func() {
			if er := recover(); er != nil {
				err = errors.Errorf("%v", er)
			}
		}()
		aDatum := args[0]
		if aDatum.Kind() == types.KindNull {
			d.SetNull()
			return
		}
		switch op {
		case opcode.Not:
			var n int64
			n, err = aDatum.ToBool()
			if err != nil {
				err = errors.Trace(err)
			} else if n == 0 {
				d.SetInt64(1)
			} else {
				d.SetInt64(0)
			}
		case opcode.BitNeg:
			var n int64
			// for bit operation, we will use int64 first, then return uint64
			n, err = aDatum.ToInt64()
			if err != nil {
				err = errors.Trace(err)
				return
			}
			d.SetUint64(uint64(^n))
		case opcode.Plus:
			switch aDatum.Kind() {
			case types.KindInt64,
				types.KindUint64,
				types.KindFloat64,
				types.KindFloat32,
				types.KindMysqlDuration,
				types.KindMysqlTime,
				types.KindString,
				types.KindMysqlDecimal,
				types.KindBytes,
				types.KindMysqlHex,
				types.KindMysqlBit,
				types.KindMysqlEnum,
				types.KindMysqlSet:
				d = aDatum
			default:
				err = ErrInvalidOperation
				return
			}
		case opcode.Minus:
			switch aDatum.Kind() {
			case types.KindInt64:
				d.SetInt64(-aDatum.GetInt64())
			case types.KindUint64:
				d.SetInt64(-int64(aDatum.GetUint64()))
			case types.KindFloat64:
				d.SetFloat64(-aDatum.GetFloat64())
			case types.KindFloat32:
				d.SetFloat32(-aDatum.GetFloat32())
			case types.KindMysqlDuration:
				d.SetMysqlDecimal(mysql.ZeroDecimal.Sub(aDatum.GetMysqlDuration().ToNumber()))
			case types.KindMysqlTime:
				d.SetMysqlDecimal(mysql.ZeroDecimal.Sub(aDatum.GetMysqlTime().ToNumber()))
			case types.KindString, types.KindBytes:
				f, err := types.StrToFloat(aDatum.GetString())
				err = errors.Trace(err)
				d.SetFloat64(-f)
			case types.KindMysqlDecimal:
				f, _ := aDatum.GetMysqlDecimal().Float64()
				d.SetMysqlDecimal(mysql.NewDecimalFromFloat(-f))
			case types.KindMysqlHex:
				d.SetFloat64(-aDatum.GetMysqlHex().ToNumber())
			case types.KindMysqlBit:
				d.SetFloat64(-aDatum.GetMysqlBit().ToNumber())
			case types.KindMysqlEnum:
				d.SetFloat64(-aDatum.GetMysqlEnum().ToNumber())
			case types.KindMysqlSet:
				d.SetFloat64(-aDatum.GetMysqlSet().ToNumber())
			default:
				err = ErrInvalidOperation
				return
			}
		default:
			err = ErrInvalidOperation
			return
		}
		return
	}
}
