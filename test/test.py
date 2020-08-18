def zero(x):
    x = 0
    return x

def main():
    x = 5
    print(id(x))
    x = zero(x) # x will become 0
    # zero(x) # x remains as 5
    print(id(x))
    print(x)

main()