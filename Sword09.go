package main

//type CQueue struct {
//	stack1, stack2 *list.List
//}
//
//
//func Constructor() CQueue {
//	return CQueue{
//		stack1: list.New(),
//		stack2: list.New(),
//	}
//}
//
//
//func (this *CQueue) AppendTail(value int)  {
//	this.stack1.PushBack(value)
//}
//
//
//func (this *CQueue) DeleteHead() int {
//	if this.stack2.Len() == 0 {
//		for this.stack1.Len() > 0 {
//			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
//		}
//	}
//	if this.stack2.Len() != 0 {
//		e := this.stack2.Back()
//		this.stack2.Remove(e)
//		return e.Value.(int)
//	}
//	return -1
//}

type CQueue struct {
	// 用切片代替
	// 栈 A，用于添加元素
	stackA []int
	// 栈 B，用于取出元素
	stackB []int
}

// CQueue 的构造函数
func Constructor() CQueue {
	// 返回一个新的 CQueue
	return CQueue{}
}
// CQueue指针地址，函数AppendTail，形式参数value, 函数返回值类型为空void
// 往队列插入整数
func (this *CQueue) AppendTail(value int)  {
	// 在 stackA 的末尾追加 value
	this.stackA = append(this.stackA, value)
}
// CQueue指针地址，函数DeleteHead，无形式参数, 函数返回值类型为int
// 从队列取出整数并删除
func (this *CQueue) DeleteHead() int {
	// 如果 stackB 没有元素则从 stackA 中取出所有
	if len(this.stackB) == 0 {
		// 如果 stackA 里也没有元素，则队列爲空返回 -1
		if len(this.stackA) == 0 {
			return -1
		}
		// 将 stackA 的所有元素转移到 stackB
		for len(this.stackA) > 0 {
			// 获取 stackA 最末尾元素的下标
			index := len(this.stackA) - 1
			// 获取 stackA 最末尾元素的值 value
			value := this.stackA[index]
			// 向 stackB 的末尾追加 value
			this.stackB = append(this.stackB, value)
			// 从 stackA 中裁剪出末尾元素，相当去remove最后一个元素
			this.stackA = this.stackA[:index]
		}
	}
	// 这时候表示 stackB 内已有元素
	// 获取 stackB 最末尾元素的下标
	index := len(this.stackB) - 1
	// 获取 stackB 最末尾元素的值 value
	value := this.stackB[index]
	// 从 stackB 中裁剪出末尾元素
	this.stackB = this.stackB[:index]
	// 返回 value
	return value
}