from typing import Any

class Stack():
    def __init__(self):
        self.array = []
        self.size  = 0

    def __len__(self)->int:
        return self.size
        
    def Push(self, data: Any):
        self.size +=1
        self.array.append(data)
    
    def Pop(self)->Any:
        if self.IsEmpty():
            return None
        data = self.array[0]
        self.array.remove(self.array[0])
        self.size -= 1
        
        return data

    def Reset(self):
        self.array = None
        self.size = 0

    def IsEmpty(self)->bool:
        return self.size == 0

if __name__ == "__main__":
    stack = Stack()
    stack.Push(0)
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    stack.Push(4)
    stack.Push(5)
    print("stack size:%d\n" %(len(stack)))
    while not stack.IsEmpty():
        print(stack.Pop())
        

        
