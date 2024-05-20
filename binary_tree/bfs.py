from collections import deque

class TreeNode:
    def __init__(self, key):
        self.left = None
        self.right = None
        self.val = key

def bfs(root):
    q = deque([root])
    res = 0

    while q:
        node = q.popleft() #dequeue
        if node:
            res +=1
        q.append(node.left) #enqueue left child
        q.append(node.right) #enqueue right child
    return res


if __name__ == "__main__":
    bt = TreeNode()
    bfs(bt)