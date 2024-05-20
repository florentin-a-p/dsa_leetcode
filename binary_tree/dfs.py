def dfs(node):
    if not node:
        return 0

    lval = dfs(node.left)
    rval = dfs(node.right)

    toadd = 1 if node.val %2==0 else 0
    return  toadd + lval + rval