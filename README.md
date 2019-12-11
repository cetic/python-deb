### Python-deb
This project explore the possibility to convert a python project to .deb package
in gitlab CI pipeline. The installation of the built package will be into a virtual environment.

### Requirements
A python project which respect the following architecture :

```
+-- project_folder
|   +-- __init__.py
+-- setup.py
+-- MANIFEST.in
```

### Usage
Just put the `.gitlab-ci.yaml` a the root of your folder. And merge your working branch with a branch called `develop`.
After that Gitlab do the rest.

### How does that work?
The Gitlab Ci use the [selltom/python-deb:v4](https://hub.docker.com/repository/docker/selltom/python-deb) docker image which contains some required tools:

#### 1. setup2control
this tool is just a parser/mapper from the `setup.py` file to the `control` file needed to build a debian package.

For example this `setup.py` file :
```
from setuptools import setup

setup(
        name='my_awesome_python_project',
        version='0.0.1',
        author='John Doe',
        description='Sample project',
        long_description='This project is just a example used to explain the python-dev CI pipeline',
        url='',
        author_email='john@doe.com',
        license='Proprietary',
        packages=[],
        include_package_data=True,
        install_requires=[],
        zip_safe=False,
        entry_points={
                'console_scripts': ['myapp = my_awesome_python_project']
        }
)

```

will create the following `control` file:

```
Architecture: amd64
Essential: no
Priority: optional
Depends: python3-pip, python3-venv, systemd, mosquitto
Package: my_awesome_python_project
Version: 0.0.1
Maintainer: John Doe
Description: This project is just a example used to explain the python-dev CI pipeline
```
