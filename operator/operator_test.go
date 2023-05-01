package operator

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	// math: + - * / % ++ --
	// relation: == != < <= > >=
	// logic && || !
	// bitwise: & | ^ << >>
	/*
		=	简单的赋值运算符，将一个表达式的值赋给一个左值	C = A + B 将 A + B 表达式结果赋值给 C
		+=	相加后再赋值	C += A 等于 C = C + A
		-=	相减后再赋值	C -= A 等于 C = C - A
		*=	相乘后再赋值	C *= A 等于 C = C * A
		/=	相除后再赋值	C /= A 等于 C = C / A
		%=	求余后再赋值	C %= A 等于 C = C % A
		<<=	左移后赋值	C <<= 2 等于 C = C << 2
		>>=	右移后赋值	C >>= 2 等于 C = C >> 2
		&=	按位与后赋值	C &= 2 等于 C = C & 2
		^=	按位异或后赋值	C ^= 2 等于 C = C ^ 2
		|=	按位或后赋值	C |= 2 等于 C = C | 2
	*/
	// other : & *
	// pointer variable can only be used as a pointer, not a value for compute
	var a int = 10
	var b *int = &a
	*b = 20
	fmt.Println(*b)
}
