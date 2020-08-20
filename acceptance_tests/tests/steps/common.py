import requests

from behave import step
from behave.runner import Context


@step('I make a "{method}" request to "{url}"')
def make_a_request(context: Context, method: str, url: str):
	req = getattr(requests, method.lower())
	req(url)