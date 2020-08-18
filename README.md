# etcd-cron-job
This package has been based on the project https://github.com/robfig/cron

# purpose

This package aims at implementing a distributed and fault tolerant cron in order to:

* Run an identical process on several hosts
* Each of these process instantiate a cron with the same rules
* Ensure only one of these processes executes an iteration of a job
