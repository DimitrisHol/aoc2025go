import os
import z3

def z3Solver(buttons, goal) : 

    presses = []
    for i in range(len(buttons)) : 
        presses.append(z3.Int(i))

    # Optimize since we want the min of presses
    s = z3.Optimize()

    # Number of press is a positive integer
    for press in presses : 
        s.add(z3.And(press >= 0))

    # The presses must equal to the goal
    # Basically create the matrix

    # We need to add the rows (the goals)
    for goalIndex, goalValue in enumerate(goal) : 

        total = 0

        # Check which presses affect this goal 
        for buttonIndex, button in enumerate(buttons) : 

            if goalIndex in button : 
                # Total transcends integer and becomes something like 
                # presses[4] + presses[5]
                total += presses[buttonIndex]

        # And this basically creates the presses[4] + presses[5] = 5
        # This is the requirement we need to add
        s.add(total == goalValue)
   

    s.minimize(sum(presses))
    assert s.check() == z3.sat

    m = s.model()
    return sum(m[press].as_long() for press in presses)

with open("../../input/10.txt", "r") as sourceFile : 


    part2 = 0

    for line in sourceFile : 

        factory = line.strip().split(" ")
        
        wiringSchematics = factory[1 : len(factory)-1]    # e.g. [(3), (1,3), (2)] etc.
        joltageRequirements = factory[len(factory)-1:][0] # e.g. {3,5,4,7}

        # insane
        intButtons = [[int(num) for num in x[1:-1].split(",")] for x in wiringSchematics]
        intGoal = [int(num) for num in joltageRequirements[1:len(joltageRequirements)- 1].split(",")]


        part2 += z3Solver(intButtons, intGoal)

    print(part2)
