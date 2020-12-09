converttosimdata.py integrates VRUS from terminal input and csv file and converts them into a json file. Current script parses the udacity open data set: https://github.com/udacity/self-driving-car/tree/master/datasets


## TERMINAL INPUT GUIDE
* Road user types in order: BIKE, PEDESTRIAN, CAR, TRUCK, SKATER
* Argument 0 --- converttosimdata.py
* Argument 1 --- csv file that contains two columns "lat" and "long", both in float or integer.
* Argument 2 --- The name for the json file. Must contain ".json". Optional.
* Argument 3 --- Number of routes. Must be a positive integer.
* Argument 4 --- List of BIKE road users. If there are no bikers but other following road users, input a list of zeros. If there are no following road users, optional. Length of list must be equal to number of routes. Elements in the list must be positive integers.
* Argument 5 --- List of PEDESTRIAN road users. If there are no pedestrians but other following road users, input a list of zeros. If there are no pedestrians and no following road users, no input is needed.
* Argument 6 --- List of CAR road users.
* Argument 7 --- List of TRUCK road users.
* Argument 8 --- List of SKATER road users.
* Arguments not noted optional are required.

## TO RUN
1. Download converttosimdata.py.
2. From terminal, run for example:
  * python converttosimdata.py gps.csv output1.json 2 [2,2] [3,3] [4,4] [5,5] [6,6]
  * python converttosimdata.py gps.csv output2.json 2 [2,2] [0,0] [4,4]
  * python converttosimdata.py gps.csv output3.json 3 [0,0,0] [3,3,0] [0,0,0] [1,0,1]
