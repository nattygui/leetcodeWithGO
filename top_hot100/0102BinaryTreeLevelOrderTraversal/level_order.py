# coding: utf-8

class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        result = []
        if root == None:
            return result
        
        cur_level_nodes = [root]
        while True:
            if len(cur_level_nodes) == 0:
                break
            next_level_nodes = []
            cur_level_values = []
            for node in cur_level_nodes:
                cur_level_values.append(node.val)
                if node.left != None:
                    next_level_nodes.append(node.left)
                if node.right != None:
                    next_level_nodes.append(node.right)

            result.append(cur_level_values)
            cur_level_nodes = next_level_nodes
        return result