# Noise: container to test the predictability of the binary execution

This simple program continually compares the time after sleeping for a second and plot a dot if the absolute value of difference between two executions is greater that a treshold.
Every minute it changes the line, so you can see how often during a minute its execution did not happened periodically.


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
