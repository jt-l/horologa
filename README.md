Horologa
=

Determine the utilization, hyper period, and frame length to be used in a static clock-driven scheduler. 

## Utilization

This program calculates the rate-monotonic utilization of a set of tasks.

Liu & Layland (1973) proved that for a set of n period tasks with unique periods, a feasible schedule that will always meet deadlines exists if the CPU utilization is below a specific bound (depending on the number of tasks). The scheduling test for this is the following equation: 

Utility = sum i = 1 to n [e_i/c_i] <= (2^(1/n) - 1)

## Hyper Period

The hyper period for a set of tasks is the least common multiple of all periods of the given tasks. The hyper period gives the minimum amount of time that the schedule needs to be calculated for before the tasks can just repeat (on the determined schedule over the hyper period).

## Frame Length

Given a set of tasks where each task contains a phase, period, execution time, and relative deadline horologa will determine if there is a valid frame length that can be used for the provided set of tasks. 

Static clock-driven schedulers are applicable to deterministic systems in which the parameters of all periodic tasks are known a priori. A static schedule can be computed in advance. 

To simplify implementation of these shedulers often it is recommended to impose structure, this is done by introducing periodic intervals called frames. Scheduling decisions are to be made at each frame. This gives two main benefits; the scheduler can check for overruns/missed deadlines at the end of each frame and the scheduler can use a periodic clock interrupt rather than a programable timer.

This program determines a valid frame length by searching for frames which statisfy all three of the following constraints: 

1. frame length >= max(e_1, e_2, ... e_n)
2. There must exist some p_i such that frame length % p_i == 0
3. 2*frame length - gcd(p_i, frame length) > D_i for i = 1 ... n

If the program cannot find a valid frame length this means that the given tasks cannot satisfy all three constraints at once. In this case you can attempt to generate new tasks by paritioning a task with large execution time into smaller sub tasks.

