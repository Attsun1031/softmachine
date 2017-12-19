# -*- coding: utf-8 -*-
import pathlib
import json
import pytest
from jsonschema import validate, ValidationError


class TestWorkflowSchema:
    with pathlib.Path(__file__).parent.joinpath("workflow-schema.json").open() as f:
        schema = json.load(f)

    def test_empty(self):
        with pytest.raises(ValidationError):
            validate({}, self.schema)

    def test_empty_tasks(self):
        with pytest.raises(ValidationError):
            validate({
                "name": "test_valid",
                "tasks": []
            }, self.schema)

    def test_single_job(self):
        validate({
            "name": "test_valid",
            "tasks": [{
                "name": "kube-job-task",
                "type": "kube-job",
                "job": {
                    "metadata": {
                        "name": "pi"
                    },
                    "spec": {
                        "template": {
                            "spec": {
                                "containers": [
                                    {
                                        "name": "pi",
                                        "command": ["perl", "print bpi"]
                                    }
                                ]
                            }
                        }
                    }
                }
            }]
        }, self.schema)

    def test_parallel(self):
        validate({
            "name": "test_valid",
            "tasks": [
                {
                    "name": "task1",
                    "type": "kube-job",
                    "job": {"x": 1},
                    "next": "para1"
                },
                {
                    "name": "para1",
                    "type": "parallel",
                    "next": "task4",
                    "tasks": [
                        {
                            "name": "task2",
                            "type": "kube-job",
                            "job": {"x": 2},
                        },
                        {
                            "name": "task3",
                            "type": "kube-job",
                            "job": {"x": 3},
                        }
                    ]
                },
                {
                    "name": "task4",
                    "type": "kube-job",
                    "job": {"x": 4},
                },
            ]
        }, self.schema)

    def test_branch(self):
        validate({
            "name": "test_valid",
            "tasks": [
                {
                    "name": "task1",
                    "type": "kube-job",
                    "job": {"y": "hoge"},
                    "next": "para1"
                },
                {
                    "name": "branch1",
                    "type": "branch",
                    "tasks": [
                        {
                            "condition": "$.value > 3",
                            "next": "task2"
                        },
                        {
                            "condition": "$.value < 3",
                            "next": "task3"
                        }
                    ]
                },
                {
                    "name": "task2",
                    "type": "kube-job",
                    "job": {"y": "hoge2"},
                },
                {
                    "name": "task3",
                    "type": "kube-job",
                    "job": {"y": "hoge3"},
                }
            ]
        }, self.schema)
