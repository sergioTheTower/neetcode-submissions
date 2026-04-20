class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:	
        rows = defaultdict(set)
        cols = defaultdict(set)
        squares = defaultdict(set)

        for r in range(9):
            for c in range(9):
                board_value = board[r][c]
                if board_value == ".":
                    continue

                if (
                    board_value in rows[r]
                    or board_value in cols[c]
                    or board_value in squares[(r // 3, c // 3)]
                ):
                    return False
                rows[r].add(board_value)
                cols[c].add(board_value)
                squares[(r // 3, c // 3)].add(board_value)

        return True

        