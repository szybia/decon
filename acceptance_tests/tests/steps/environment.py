import time

import requests

def before_all(context):
	#	wait for decon to come up for 30 seconds
	for _ in range(30 * 4):
		try:
			requests.get('http://decon/health')
			break
		except requests.exceptions.ConnectionError:
			time.speep(0.25)
	assert False, 'decon failed to come up after 30 seconds'