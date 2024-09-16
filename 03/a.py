import gpiozero
import time

sw = gpiozero.DigitalInputDevice(9,pull_up = True)

count = 0

while True:
	if(sw.value == 0):
		count = count + 1
		print("Count: " + str(count))
		while(sw.value == 0):
			time.sleep(0.1)
 
