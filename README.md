# Noise: container to test the predictability of the binary execution

This simple program continually compares the time after sleeping for a second and plot a dot if the absolute value of the difference between two executions is greater than a threshold.
It uses a similar approach of: https://docs.azul.com/prime/jHiccup
Every minute it starts a new line, so you can see how often during a minute its execution did not happened periodically.


```
kubectl logs -f noise
......
....
.....
......
.........
.....
.......
......
.....
....
.....
.........
.
....
......

``` 

in the above example you can see during the first minute of its execution, for 7 times the interval between executions was far away from a second: plus or minus 0.5 milliseconds

Now you can compare settings in your K8s cluster to try to make it more predictable 
