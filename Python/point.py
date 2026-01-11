import math
# class point:
#     x = 0
#     y = 0

#     def set_location(self, a, b):
#         self.x = a
#         self.y = b

#     def translate(self, dx, dy):
#         self.x += dx
#         self.y += dy

#     def distance(self, x1, y1):
#         d = math.sqrt((self.x - x1)**2+(self.y - y1)**2)
#         return d

#     def distance_from_orogin(self):
#         d = math.sqrt((self.x)**2 + (self.y)**2)
#         return d

# p1 = point()
# p1.x = 2
# p1.y = -5
# print(p1.x)
# print(p1.y)

# p1.translate(-5,7)
# print(p1.x)
# print(p1.y)

# print(p1.distance(10,10))
# print(p1.distance_from_orogin)

class point:
    x = 0
    y = 0

    def __init__(self, a, b):
        self.x = a
        self.y = b

    def translate(self, dx, dy):
        self.x += dx
        self.y += dy

    def distance(self, x1, y1):
        d = math.sqrt((self.x - other X1)**2+(self.y - y1)**2)
        return d

    def distance_from_orogin(self):
        d = math.sqrt((self.x)**2 + (self.y)**2)
        return d

def main():
    p1 = point(2,-5)
    p2 = point(3,9)

    print(p1.x)
    print(p1.y)
    print(p2.x)
    print(p2.y)

    p1.translate(-5,7)
    print(p1.x)
    print(p1.y)

    print(p1.distance(P2))
    print(p1.distance_from_orogin())

