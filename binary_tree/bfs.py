from collections import deque

class TreeNode:
    def __init__(self, key):
        self.left = None
        self.right = None
        self.val = key

class BinaryTree:
    def __init__(self):
        self.root = None

    def insert(self, key):
        if self.root is None:
            self.root = TreeNode(key)
        else:
            self._insert(self.root, key)

    def _insert(self, node, key):
        if key < node.val:
            if node.left is None:
                node.left = TreeNode(key)
            else:
                self._insert(node.left, key)
        else:
            if node.right is None:
                node.right = TreeNode(key)
            else:
                self._insert(node.right, key)

    def bfs_count_nodes(self, root):
        if not root:
            return 0

        q = deque([root]) #initialize queue
        # print("q.count: ", q.count(q))
        res = 0

        while q:
            print("Current queue: ", [node.val for node in q])

            node = q.popleft() # dequeue, because the node is being visited
            print("node: ", node.val)

            if node:
                res += 1 # the counter
                if node.left:
                    print("node.left: ", node.left.val)
                    q.append(node.left) # enqueue left child, we're going to visit this
                if node.right:
                    print("node.right: ", node.right.val)
                    q.append(node.right) # enqueue right child, we're going to visit this
            # print("\n")

        return res

    def bfs_sum_nodes(self, root):
        if not root:
            return 0

        q = deque([root]) #initialize queue
        res = 0

        while q:
            node = q.popleft() # dequeue, because the node is being visited

            if node:
                res = res + node.val # the counter
                if node.left:
                    q.append(node.left) # enqueue left child, we're going to visit this
                if node.right:
                    q.append(node.right) # enqueue right child, we're going to visit this
        return res

    def bfs_sum_odd_nodes(self, root):
        if not root:
            return 0

        q = deque([root]) #initialize queue
        res = 0

        while q:
            node = q.popleft() # dequeue, because the node is being visited

            if node:
                if node.val % 2 == 1:
                    res = res + node.val # the counter
                if node.left:
                    q.append(node.left) # enqueue left child, we're going to visit this
                if node.right:
                    q.append(node.right) # enqueue right child, we're going to visit this
        return res

    def bfs_sum_even_nodes(self, root):
        if not root:
            return 0

        q = deque([root]) #initialize queue
        res = 0

        while q:
            node = q.popleft() # dequeue, because the node is being visited

            if node:
                if node.val % 2 == 0:
                    res = res + node.val # the counter
                if node.left:
                    q.append(node.left) # enqueue left child, we're going to visit this
                if node.right:
                    q.append(node.right) # enqueue right child, we're going to visit this
        return res


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0

        q = deque([root]) #initialize queue
        level = 0

        while q:
            node = q.popleft() # dequeue, because the node is being visited

            if node:
                if node.left:
                    q.append(node.left) # enqueue left child, we're going to visit this
                if node.right:
                    q.append(node.right) # enqueue right child, we're going to visit this

            level += 1

        return level

if __name__ == "__main__":
    bt = BinaryTree()
    elements = [10, 5, 20, 3, 2, 7, 15, 1, 25, 30, 31, 32, 33]
    for elem in elements:
        bt.insert(elem)

    print("count of all nodes: ", bt.bfs_count_nodes(bt.root))
    # print("sum of all nodes: ", bt.bfs_sum_nodes(bt.root))
    # print("sum of all odd nodes: ", bt.bfs_sum_odd_nodes(bt.root))
    # print("sum of all even nodes: ", bt.bfs_sum_even_nodes(bt.root))





