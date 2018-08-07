# -*- coding: utf-8 -*-

# Learn more: https://github.com/kennethreitz/setup.py

from setuptools import setup, find_packages

setup(
    name='EasemlSchema',
    version='0.1.0',
    description='Schema which is used to define the type of a machine learning data set.',
    #long_description=readme,
    author='Bojan Karlas',
    author_email='bojan.karlas@gmail.com',
    url='https://github.com/DS3Lab/easeml-schema',
    license='MIT',
    install_requires=['numpy>=1.9.1'],
    packages=find_packages()
)
