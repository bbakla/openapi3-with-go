# CreateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int64** |  | 
**ListName** | **string** |  | 
**ListType** | [**ListType**](ListType.md) |  | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewCreateList

`func NewCreateList(userId int64, listName string, listType ListType, ) *CreateList`

NewCreateList instantiates a new CreateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListWithDefaults

`func NewCreateListWithDefaults() *CreateList`

NewCreateListWithDefaults instantiates a new CreateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateList) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateList) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateList) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetListName

`func (o *CreateList) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *CreateList) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *CreateList) SetListName(v string)`

SetListName sets ListName field to given value.


### GetListType

`func (o *CreateList) GetListType() ListType`

GetListType returns the ListType field if non-nil, zero value otherwise.

### GetListTypeOk

`func (o *CreateList) GetListTypeOk() (*ListType, bool)`

GetListTypeOk returns a tuple with the ListType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListType

`func (o *CreateList) SetListType(v ListType)`

SetListType sets ListType field to given value.


### GetTags

`func (o *CreateList) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateList) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateList) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateList) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


