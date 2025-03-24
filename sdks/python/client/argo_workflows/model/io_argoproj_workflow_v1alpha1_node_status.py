"""
    Argo Workflows API

<<<<<<< HEAD
    Argo Workflows is an open source container-native workflow engine for orchestrating parallel jobs on Kubernetes. For more information, please see https://argo-workflows.readthedocs.io/en/release-3.5/  # noqa: E501
=======
    Argo Workflows is an open source container-native workflow engine for orchestrating parallel jobs on Kubernetes. For more information, please see https://argo-workflows.readthedocs.io/en/latest/  # noqa: E501
>>>>>>> draft-3.6.5

    The version of the OpenAPI document: VERSION
    Generated by: https://openapi-generator.tech
"""


import re  # noqa: F401
import sys  # noqa: F401

from argo_workflows.model_utils import (  # noqa: F401
    ApiTypeError,
    ModelComposed,
    ModelNormal,
    ModelSimple,
    cached_property,
    change_keys_js_to_python,
    convert_js_args_to_python_args,
    date,
    datetime,
    file_type,
    none_type,
    validate_get_composed_info,
    OpenApiModel
)
from argo_workflows.exceptions import ApiAttributeError


def lazy_import():
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_inputs import IoArgoprojWorkflowV1alpha1Inputs
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_memoization_status import IoArgoprojWorkflowV1alpha1MemoizationStatus
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_node_flag import IoArgoprojWorkflowV1alpha1NodeFlag
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_node_synchronization_status import IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_outputs import IoArgoprojWorkflowV1alpha1Outputs
    from argo_workflows.model.io_argoproj_workflow_v1alpha1_template_ref import IoArgoprojWorkflowV1alpha1TemplateRef
    globals()['IoArgoprojWorkflowV1alpha1Inputs'] = IoArgoprojWorkflowV1alpha1Inputs
    globals()['IoArgoprojWorkflowV1alpha1MemoizationStatus'] = IoArgoprojWorkflowV1alpha1MemoizationStatus
    globals()['IoArgoprojWorkflowV1alpha1NodeFlag'] = IoArgoprojWorkflowV1alpha1NodeFlag
    globals()['IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus'] = IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus
    globals()['IoArgoprojWorkflowV1alpha1Outputs'] = IoArgoprojWorkflowV1alpha1Outputs
    globals()['IoArgoprojWorkflowV1alpha1TemplateRef'] = IoArgoprojWorkflowV1alpha1TemplateRef


class IoArgoprojWorkflowV1alpha1NodeStatus(ModelNormal):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.

    Attributes:
      allowed_values (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          with a capitalized key describing the allowed value and an allowed
          value. These dicts store the allowed enum values.
      attribute_map (dict): The key is attribute name
          and the value is json key in definition.
      discriminator_value_class_map (dict): A dict to go from the discriminator
          variable value to the discriminator class name.
      validations (dict): The key is the tuple path to the attribute
          and the for var_name this is (var_name,). The value is a dict
          that stores validations for max_length, min_length, max_items,
          min_items, exclusive_maximum, inclusive_maximum, exclusive_minimum,
          inclusive_minimum, and regex.
      additional_properties_type (tuple): A tuple of classes accepted
          as additional properties values.
    """

    allowed_values = {
    }

    validations = {
    }

    @cached_property
    def additional_properties_type():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded
        """
        lazy_import()
        return (bool, date, datetime, dict, float, int, list, str, none_type,)  # noqa: E501

    _nullable = False

    @cached_property
    def openapi_types():
        """
        This must be a method because a model may have properties that are
        of type self, this must run after the class is loaded

        Returns
            openapi_types (dict): The key is attribute name
                and the value is attribute type.
        """
        lazy_import()
        return {
            'id': (str,),  # noqa: E501
            'name': (str,),  # noqa: E501
            'type': (str,),  # noqa: E501
            'boundary_id': (str,),  # noqa: E501
            'children': ([str],),  # noqa: E501
            'daemoned': (bool,),  # noqa: E501
            'display_name': (str,),  # noqa: E501
            'estimated_duration': (int,),  # noqa: E501
            'finished_at': (datetime,),  # noqa: E501
            'host_node_name': (str,),  # noqa: E501
            'inputs': (IoArgoprojWorkflowV1alpha1Inputs,),  # noqa: E501
            'memoization_status': (IoArgoprojWorkflowV1alpha1MemoizationStatus,),  # noqa: E501
            'message': (str,),  # noqa: E501
            'node_flag': (IoArgoprojWorkflowV1alpha1NodeFlag,),  # noqa: E501
            'outbound_nodes': ([str],),  # noqa: E501
            'outputs': (IoArgoprojWorkflowV1alpha1Outputs,),  # noqa: E501
            'phase': (str,),  # noqa: E501
            'pod_ip': (str,),  # noqa: E501
            'progress': (str,),  # noqa: E501
            'resources_duration': ({str: (int,)},),  # noqa: E501
            'started_at': (datetime,),  # noqa: E501
            'synchronization_status': (IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus,),  # noqa: E501
            'template_name': (str,),  # noqa: E501
            'template_ref': (IoArgoprojWorkflowV1alpha1TemplateRef,),  # noqa: E501
            'template_scope': (str,),  # noqa: E501
        }

    @cached_property
    def discriminator():
        return None


    attribute_map = {
        'id': 'id',  # noqa: E501
        'name': 'name',  # noqa: E501
        'type': 'type',  # noqa: E501
        'boundary_id': 'boundaryID',  # noqa: E501
        'children': 'children',  # noqa: E501
        'daemoned': 'daemoned',  # noqa: E501
        'display_name': 'displayName',  # noqa: E501
        'estimated_duration': 'estimatedDuration',  # noqa: E501
        'finished_at': 'finishedAt',  # noqa: E501
        'host_node_name': 'hostNodeName',  # noqa: E501
        'inputs': 'inputs',  # noqa: E501
        'memoization_status': 'memoizationStatus',  # noqa: E501
        'message': 'message',  # noqa: E501
        'node_flag': 'nodeFlag',  # noqa: E501
        'outbound_nodes': 'outboundNodes',  # noqa: E501
        'outputs': 'outputs',  # noqa: E501
        'phase': 'phase',  # noqa: E501
        'pod_ip': 'podIP',  # noqa: E501
        'progress': 'progress',  # noqa: E501
        'resources_duration': 'resourcesDuration',  # noqa: E501
        'started_at': 'startedAt',  # noqa: E501
        'synchronization_status': 'synchronizationStatus',  # noqa: E501
        'template_name': 'templateName',  # noqa: E501
        'template_ref': 'templateRef',  # noqa: E501
        'template_scope': 'templateScope',  # noqa: E501
    }

    read_only_vars = {
    }

    _composed_schemas = {}

    @classmethod
    @convert_js_args_to_python_args
    def _from_openapi_data(cls, id, name, type, *args, **kwargs):  # noqa: E501
        """IoArgoprojWorkflowV1alpha1NodeStatus - a model defined in OpenAPI

        Args:
            id (str): ID is a unique identifier of a node within the worklow It is implemented as a hash of the node name, which makes the ID deterministic
            name (str): Name is unique name in the node tree used to generate the node ID
            type (str): Type indicates type of node

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            boundary_id (str): BoundaryID indicates the node ID of the associated template root node in which this node belongs to. [optional]  # noqa: E501
            children ([str]): Children is a list of child node IDs. [optional]  # noqa: E501
            daemoned (bool): Daemoned tracks whether or not this node was daemoned and need to be terminated. [optional]  # noqa: E501
            display_name (str): DisplayName is a human readable representation of the node. Unique within a template boundary. [optional]  # noqa: E501
            estimated_duration (int): EstimatedDuration in seconds.. [optional]  # noqa: E501
            finished_at (datetime): Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.. [optional]  # noqa: E501
            host_node_name (str): HostNodeName name of the Kubernetes node on which the Pod is running, if applicable. [optional]  # noqa: E501
            inputs (IoArgoprojWorkflowV1alpha1Inputs): [optional]  # noqa: E501
            memoization_status (IoArgoprojWorkflowV1alpha1MemoizationStatus): [optional]  # noqa: E501
            message (str): A human readable message indicating details about why the node is in this condition.. [optional]  # noqa: E501
            node_flag (IoArgoprojWorkflowV1alpha1NodeFlag): [optional]  # noqa: E501
            outbound_nodes ([str]): OutboundNodes tracks the node IDs which are considered \"outbound\" nodes to a template invocation. For every invocation of a template, there are nodes which we considered as \"outbound\". Essentially, these are last nodes in the execution sequence to run, before the template is considered completed. These nodes are then connected as parents to a following step.  In the case of single pod steps (i.e. container, script, resource templates), this list will be nil since the pod itself is already considered the \"outbound\" node. In the case of DAGs, outbound nodes are the \"target\" tasks (tasks with no children). In the case of steps, outbound nodes are all the containers involved in the last step group. NOTE: since templates are composable, the list of outbound nodes are carried upwards when a DAG/steps template invokes another DAG/steps template. In other words, the outbound nodes of a template, will be a superset of the outbound nodes of its last children.. [optional]  # noqa: E501
            outputs (IoArgoprojWorkflowV1alpha1Outputs): [optional]  # noqa: E501
            phase (str): Phase a simple, high-level summary of where the node is in its lifecycle. Can be used as a state machine. Will be one of these values \"Pending\", \"Running\" before the node is completed, or \"Succeeded\", \"Skipped\", \"Failed\", \"Error\", or \"Omitted\" as a final state.. [optional]  # noqa: E501
            pod_ip (str): PodIP captures the IP of the pod for daemoned steps. [optional]  # noqa: E501
            progress (str): Progress to completion. [optional]  # noqa: E501
            resources_duration ({str: (int,)}): ResourcesDuration is indicative, but not accurate, resource duration. This is populated when the nodes completes.. [optional]  # noqa: E501
            started_at (datetime): Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.. [optional]  # noqa: E501
            synchronization_status (IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus): [optional]  # noqa: E501
            template_name (str): TemplateName is the template name which this node corresponds to. Not applicable to virtual nodes (e.g. Retry, StepGroup). [optional]  # noqa: E501
            template_ref (IoArgoprojWorkflowV1alpha1TemplateRef): [optional]  # noqa: E501
            template_scope (str): TemplateScope is the template scope in which the template of this node was retrieved.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        self = super(OpenApiModel, cls).__new__(cls)

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.id = id
        self.name = name
        self.type = type
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
        return self

    required_properties = set([
        '_data_store',
        '_check_type',
        '_spec_property_naming',
        '_path_to_item',
        '_configuration',
        '_visited_composed_classes',
    ])

    @convert_js_args_to_python_args
    def __init__(self, id, name, type, *args, **kwargs):  # noqa: E501
        """IoArgoprojWorkflowV1alpha1NodeStatus - a model defined in OpenAPI

        Args:
            id (str): ID is a unique identifier of a node within the worklow It is implemented as a hash of the node name, which makes the ID deterministic
            name (str): Name is unique name in the node tree used to generate the node ID
            type (str): Type indicates type of node

        Keyword Args:
            _check_type (bool): if True, values for parameters in openapi_types
                                will be type checked and a TypeError will be
                                raised if the wrong type is input.
                                Defaults to True
            _path_to_item (tuple/list): This is a list of keys or values to
                                drill down to the model in received_data
                                when deserializing a response
            _spec_property_naming (bool): True if the variable names in the input data
                                are serialized names, as specified in the OpenAPI document.
                                False if the variable names in the input data
                                are pythonic names, e.g. snake case (default)
            _configuration (Configuration): the instance to use when
                                deserializing a file_type parameter.
                                If passed, type conversion is attempted
                                If omitted no type conversion is done.
            _visited_composed_classes (tuple): This stores a tuple of
                                classes that we have traveled through so that
                                if we see that class again we will not use its
                                discriminator again.
                                When traveling through a discriminator, the
                                composed schema that is
                                is traveled through is added to this set.
                                For example if Animal has a discriminator
                                petType and we pass in "Dog", and the class Dog
                                allOf includes Animal, we move through Animal
                                once using the discriminator, and pick Dog.
                                Then in Dog, we will make an instance of the
                                Animal class but this time we won't travel
                                through its discriminator because we passed in
                                _visited_composed_classes = (Animal,)
            boundary_id (str): BoundaryID indicates the node ID of the associated template root node in which this node belongs to. [optional]  # noqa: E501
            children ([str]): Children is a list of child node IDs. [optional]  # noqa: E501
            daemoned (bool): Daemoned tracks whether or not this node was daemoned and need to be terminated. [optional]  # noqa: E501
            display_name (str): DisplayName is a human readable representation of the node. Unique within a template boundary. [optional]  # noqa: E501
            estimated_duration (int): EstimatedDuration in seconds.. [optional]  # noqa: E501
            finished_at (datetime): Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.. [optional]  # noqa: E501
            host_node_name (str): HostNodeName name of the Kubernetes node on which the Pod is running, if applicable. [optional]  # noqa: E501
            inputs (IoArgoprojWorkflowV1alpha1Inputs): [optional]  # noqa: E501
            memoization_status (IoArgoprojWorkflowV1alpha1MemoizationStatus): [optional]  # noqa: E501
            message (str): A human readable message indicating details about why the node is in this condition.. [optional]  # noqa: E501
            node_flag (IoArgoprojWorkflowV1alpha1NodeFlag): [optional]  # noqa: E501
            outbound_nodes ([str]): OutboundNodes tracks the node IDs which are considered \"outbound\" nodes to a template invocation. For every invocation of a template, there are nodes which we considered as \"outbound\". Essentially, these are last nodes in the execution sequence to run, before the template is considered completed. These nodes are then connected as parents to a following step.  In the case of single pod steps (i.e. container, script, resource templates), this list will be nil since the pod itself is already considered the \"outbound\" node. In the case of DAGs, outbound nodes are the \"target\" tasks (tasks with no children). In the case of steps, outbound nodes are all the containers involved in the last step group. NOTE: since templates are composable, the list of outbound nodes are carried upwards when a DAG/steps template invokes another DAG/steps template. In other words, the outbound nodes of a template, will be a superset of the outbound nodes of its last children.. [optional]  # noqa: E501
            outputs (IoArgoprojWorkflowV1alpha1Outputs): [optional]  # noqa: E501
            phase (str): Phase a simple, high-level summary of where the node is in its lifecycle. Can be used as a state machine. Will be one of these values \"Pending\", \"Running\" before the node is completed, or \"Succeeded\", \"Skipped\", \"Failed\", \"Error\", or \"Omitted\" as a final state.. [optional]  # noqa: E501
            pod_ip (str): PodIP captures the IP of the pod for daemoned steps. [optional]  # noqa: E501
            progress (str): Progress to completion. [optional]  # noqa: E501
            resources_duration ({str: (int,)}): ResourcesDuration is indicative, but not accurate, resource duration. This is populated when the nodes completes.. [optional]  # noqa: E501
            started_at (datetime): Time is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.. [optional]  # noqa: E501
            synchronization_status (IoArgoprojWorkflowV1alpha1NodeSynchronizationStatus): [optional]  # noqa: E501
            template_name (str): TemplateName is the template name which this node corresponds to. Not applicable to virtual nodes (e.g. Retry, StepGroup). [optional]  # noqa: E501
            template_ref (IoArgoprojWorkflowV1alpha1TemplateRef): [optional]  # noqa: E501
            template_scope (str): TemplateScope is the template scope in which the template of this node was retrieved.. [optional]  # noqa: E501
        """

        _check_type = kwargs.pop('_check_type', True)
        _spec_property_naming = kwargs.pop('_spec_property_naming', False)
        _path_to_item = kwargs.pop('_path_to_item', ())
        _configuration = kwargs.pop('_configuration', None)
        _visited_composed_classes = kwargs.pop('_visited_composed_classes', ())

        if args:
            raise ApiTypeError(
                "Invalid positional arguments=%s passed to %s. Remove those invalid positional arguments." % (
                    args,
                    self.__class__.__name__,
                ),
                path_to_item=_path_to_item,
                valid_classes=(self.__class__,),
            )

        self._data_store = {}
        self._check_type = _check_type
        self._spec_property_naming = _spec_property_naming
        self._path_to_item = _path_to_item
        self._configuration = _configuration
        self._visited_composed_classes = _visited_composed_classes + (self.__class__,)

        self.id = id
        self.name = name
        self.type = type
        for var_name, var_value in kwargs.items():
            if var_name not in self.attribute_map and \
                        self._configuration is not None and \
                        self._configuration.discard_unknown_keys and \
                        self.additional_properties_type is None:
                # discard variable.
                continue
            setattr(self, var_name, var_value)
            if var_name in self.read_only_vars:
                raise ApiAttributeError(f"`{var_name}` is a read-only attribute. Use `from_openapi_data` to instantiate "
                                     f"class with read only attributes.")
