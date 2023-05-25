while True:
    line = input()
    if not line:
        break
    if line.count("%") > 0:
        value = int(line[line.index("(") + 1:line.index(".")])
        print("OK" if value >= 90 else "Not OK")
        exit(0)
