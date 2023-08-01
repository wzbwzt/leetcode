boxes = [1, 2, 3, 4, 5]
status = {1: True, 2: False, 3: True, 4: False, 5: True}
q = [i for i in boxes if status[i]]
print(q)

boxes = set(boxes)
print(boxes)
