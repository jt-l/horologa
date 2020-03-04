Horologa
=

Determine the frame length to be used in a static clock-driven scheduler. 

Given a set of tasks where each task contains a phase, period, execution time, and relative deadline horologa will determine if there is a valid frame length that can be used for the provided set of tasks. 

Static clock-driven schedulers are applicable to deterministic systems in which the parameters of all periodic tasks are known a prior. A static schedule can be computed in advance. 

To simplify implementation of these shedulers often it is recommended to impose structure, this is done by introducing periodic intervals called frames. Scheduling decisions are to be made at each frame. This gives two main benefits; the scheduler can check for overruns/missed deadlines at the end of each frame and the scheduler can use a periodic clock interrupt rather than a programable timer.

This program determines a valid frame length by searching frames which statisfy all three of the following constraints: 

1. frame length >= max(e_1, e_2, ... e_n)
2. There must exist some p_i such that frame length % p_i == 0
3. 2*frame length - gcd(p_i, frame length) > D_i for i = 1 ... n

If the program cannot find a valid frame length this means that the given tasks cannot satify all three constraints at once. In this case you can attempt to generate new tasks by paritioning a task with large execution time into smaller sub tasks.

