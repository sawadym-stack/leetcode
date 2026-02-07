type MyQueue struct {
    in []int
    out []int
}


func Constructor() MyQueue {
   return MyQueue{
     in : []int{},
    out : []int{},
   }
}


func (q *MyQueue) Push(x int)  {
    q.in = append(q.in,x)
}


func (q *MyQueue) Pop() int {
    q.Peek()
    val:=q.out[len(q.out)-1]
    q.out=q.out[:len(q.out)-1]
    return val
}


func (q *MyQueue) Peek() int {
    if len(q.out)==0{
        for len(q.in)>0{
            n:= len(q.in)-1
            q.out=append(q.out,q.in[n])
            q.in=q.in[:n]
        }
    }
    return q.out[len(q.out)-1]
}


func (q *MyQueue) Empty() bool {
    return len(q.in)==0 && len(q.out)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */