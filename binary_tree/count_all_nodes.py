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

    def inorder_traversal(self, root):
        if root:
            self.inorder_traversal(root.left)
            print(root.val, end=' ')
            self.inorder_traversal(root.right)

    def pre_order_traversal(self, root):
        if root:
            print(root.val, end=' ')
            self.pre_order_traversal(root.left)
            self.pre_order_traversal(root.right)

    def post_order_traversal(self, root):
        if root:
            self.post_order_traversal(root.left)
            self.post_order_traversal(root.right)
            print(root.val, end=' ')

    def dfs_count_nodes(self, root):
        if not root:
            return 0

        lval = self.dfs_count_nodes(root.left)
        rval = self.dfs_count_nodes(root.right)

        toadd = 1 if root.val else 0
        return  toadd + lval + rval

    def dfs_sum_nodes(self, root):
        if not root: #base case
            return 0

        lval = self.dfs_sum_nodes(root.left)
        # print("lval: ", lval)

        rval = self.dfs_sum_nodes(root.right)
        # print("rval: ", rval, "\n")

        toadd = root.val if root.val else 0 # what to do with the return value?
        # print("toadd: ", toadd)

        return  toadd + lval + rval

    def dfs_sum_nodes_even(self, root):
        if not root: #base case
            return 0

        lval = self.dfs_sum_nodes_even(root.left)
        # print("lval: ", lval)

        rval = self.dfs_sum_nodes_even(root.right)
        # print("rval: ", rval, "\n")

        toadd = root.val if root.val%2 == 0 else 0 # what to do with the return value?
        # print("toadd: ", toadd)

        return  toadd + lval + rval

    def dfs_sum_nodes_odd(self, root):
        if not root: #base case
            return 0

        lval = self.dfs_sum_nodes_odd(root.left)
        rval = self.dfs_sum_nodes_odd(root.right)

        toadd = root.val if root.val%2 == 1 else 0 # what to do with the return value?
        return  toadd + lval + rval

    def dfs_sum_left_nodes_with_root(self, root):
        if not root: #base case
            return 0

        lval = self.dfs_sum_left_nodes_with_root(root.left)
        print("lval: ", lval)

        toadd = root.val if root.val else 0 # what to do with the return value?
        print("toadd: ", toadd)

        return  toadd + lval

    def dfs_sum_left_nodes_without_root(self, root):
        return  1

    def dfs_max_depth(self, root):
        if not root: #base case
            return 0

        leftnode = getattr(root.left, 'val', 0) if root.left else 0
        rightnode = getattr(root.right, 'val', 0) if root.right else 0

        lval = self.dfs_max_depth(root.left)
        rval = self.dfs_max_depth(root.right)
        # print( "left node: ", leftnode ,"right node: ", rightnode)
        print("lval: ", lval, "rval: ", rval, "left node: ", leftnode, "right node: ", rightnode)

        if root.val:
            toadd = 1
            if lval > rval:
                return toadd + lval
            else:
                return toadd + rval



if __name__ == "__main__":
    bt = BinaryTree()
    # elements = [10, 5, 20, 3, 7, 15, 30]
    elements = [10, 5, 20, 3, 2, 7, 15, 1, 25, 30, 31, 32, 33]
    # elements = [10, 5, 20, 7, 15]

    for elem in elements:
        bt.insert(elem)

    # print("Inorder traversal of the binary tree:")
    # bt.inorder_traversal(bt.root)
    #
    # print("\nPreorder traversal of the binary tree:")
    # bt.pre_order_traversal(bt.root)
    #
    # print("\nPostorder traversal of the binary tree:")
    # bt.post_order_traversal(bt.root)
    #
    # print("\ncount all nodes of the binary tree:", bt.dfs_count_nodes(bt.root))
    # print("\nsum all nodes of the binary tree:", bt.dfs_sum_nodes(bt.root))
    # print("\nsum all even nodes of the binary tree:", bt.dfs_sum_nodes_even(bt.root))
    # print("\nsum all odd nodes of the binary tree:", bt.dfs_sum_nodes_odd(bt.root))
    # print("\nsum all left nodes of the binary tree:", bt.dfs_sum_left_nodes_with_root(bt.root))

    print("\nmax depth of the binary tree:", bt.dfs_max_depth(bt.root))

    print("1 and 1: ", 1 and 1)
    print("1 and 0: ", 1 and 0)
    print("1 or 1: ", 1 or 1)
    print("1 or 0: ", 1 or 0)