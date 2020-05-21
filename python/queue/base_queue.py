"""
display base queue
"""
from typing import Any
from abc import abstractmethod,ABCMeta
from typing import Any
class QueueIf(metaclass = ABCMeta):
    @abstractmethod
    def Push(self, element):
        pass
 
    @abstractmethod
    def Pop(self)->Any:
        pass

    @abstractmethod
    def Cap(self)->int:
        pass
    @abstractmethod
    def Reset(self):
        pass
    @abstractmethod
    def IsEmpty(self)->bool:
        pass


class Queue(QueueIf):
    def __init__(self):
        self.data = []
        self.count = 0
    def __len__(self)->int:
        return self.count
    def Push(self, data):
        self.count+=1
        self.data.append(data)

    def Pop(self)-> Any:
        if self.count < 1:
            return None
        data = self.data[self.count-1]
        self.count-=1
        return data

    def Reset(self): 
        self.count=0
        self.data = []

    def Cap(self)->int:
          return self.count
    
    def IsEmpty(self)->bool:
        return self.count == 0

if __name__ == "__main__":
    q = Queue()
    q.Push("Java")
    q.Push("Go")
    q.Push("C++")
    q.Push("Python")
    print(len(q))
    print(q.Cap())
    while not q.IsEmpty():
        print(q.Pop())  



    
   
        
        


