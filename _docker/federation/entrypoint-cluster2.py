#!/usr/bin/env python
# encoding: utf-8
import os
import sys
import time
import logging
import re
import socket
from ast import literal_eval
from subprocess import check_call, check_output, Popen, PIPE


def cluster_status():
    return check_output(
        "rabbitmqctl cluster_status", shell=True
    ).decode().splitlines()[0].replace(
        'Cluster status of node', ''
    ).rstrip('.').strip()


def healthcheck():
    prc = Popen("rabbitmqctl -q node_health_check 2>/dev/null",
                shell=True, stdout=PIPE)
    prc.wait()
    if prc.returncode != 0:
        return False

    return 'Health check passed' in prc.stdout.read().decode()


def check_tcp_port(host, port):
    try:
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        result = sock.connect_ex((host, port))

        return result == 0
    except:
        return False


def main():
    logging.basicConfig(level=logging.INFO)

    logging.info("Federtaion mode")
    process = Popen(['docker-entrypoint.sh', 'rabbitmq-server'])

    while not check_tcp_port('localhost', 5672):
        logging.info('Waiting local node...')
        time.sleep(1)

    while not healthcheck():
        logging.info("Waiting for...")
        time.sleep(1)

    node_name = cluster_status()
    check_call(['rabbitmqctl', 'set_parameter', 'federation-upstream', 'rabbit@rabbit1',
                '''{"uri":"amqp://guest:guest@192.168.115.48:5672/federation"}'''])
    check_call(['rabbitmqctl', 'set_policy', '--apply-to', 'exchanges', 'federation-test', 'test',
                '''{"federation-upstream-set":"all"}'''])
    process.wait()
    return exit(process.returncode)


if __name__ == '__main__':
    main()
