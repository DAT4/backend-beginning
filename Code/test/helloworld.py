import requests,json, uuid, re

def createUser():
    url = 'https://api.backend.mama.sh/user'
    ip = requests.get('http://httpbin.org/ip')
    mac = ':'.join(re.findall('..', '%012x' % uuid.getnode()))
    user = {
            'email': 'mama@mama.sh',
            'password': 'abcabc',
            'name': 'Martin Maartensson',
            'ip': ip.json()['origin'],
            'mac': mac,
            }

    res = requests.post(url,json=user)
    print(res.text)

def helloWorld():
    url = 'https://api.backend.mama.sh/'
    msg = requests.get(url)
    print(msg.text)

def listUsers():
    url = 'https://api.backend.mama.sh/user'
    msg = requests.get(url)
    print(msg.text)

listUsers()
