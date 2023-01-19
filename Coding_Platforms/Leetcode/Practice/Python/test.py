l =[1,10,2,3,4]
# if (a == sorted(a)) & (len(a) == len(set(a))):
#     print("True")
print(sorted(l)[1],l[1])
print(all(l[:i]+l[i+1:] for i in range(len(l) - 1)))
print("Hello world")