package main

func main() {
	/*
			1.k1, k2, k3的数据类型?
					k1 := vector.IntVector{}//k1的数据类型: vector.IntVector
					k2 := &vector.IntVector{}//k2的数据类型: &vector.IntVector
					k3 := new(vector.IntVector)//k3的数据类型: &vector.IntVector
					k1.Push(2)
					k2.Push(3)
					k3.Push(4)

				2.当前，这个程序可以编译并且运行良好。在不同类型的变量上 Push 都可以工作。
				Push 的文档这样描述：
				func (p *IntVector) Push(x int) Push 增加 x 到向量的末尾。
				那么接受者应当是 *IntVector 类型，为什么上面的代码（Push 语句）可以正确工作？
				above (the Push statements) work correct then?

		当 x 的方法集合包含 m，并且参数列表可以赋值给 m 的参数，方法调用 x.m() 是合法的。
		如果 x 可以被地址化，而 &x 的方法集合包含 m，x.m() 可以作为 (&x).m() 的省略写法。
	    换句话说，由于 k1 可以被地址化，而 *vector.IntVector 具有 Push 方法，
	    调用 k1.Push(2) 被 Go 转换为 (&k1).Push(2) 来使型系统愉悦
	*/
}
