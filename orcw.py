#!/usr/local/bin/python3
import pyorc
from uuid import uuid4

with open('./data.orc', 'wb') as data:
    with pyorc.Writer(data, 'struct<col0:int,col1:string,col2:string,col3:string,col4:string>') as writer:
        for idx in range(10000000):
            uuid = str(uuid4())
            writer.write((idx, uuid + '1', uuid + '2', uuid + '3', uuid + '4'))
        