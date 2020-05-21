"""
interface for queue
"""
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


