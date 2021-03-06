package customer

import(
    "github.com/gin-gonic/gin"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/fecshopsoft/fec-go/util"
    "strconv"
    "log"
  //  "errors"
    "github.com/fecshopsoft/fec-go/helper"
)
// role_id 和resource_id 关系对应表。
type RoleResource struct {
    Id int64 `form:"id" json:"id"`
    RoleId int64 `form:"role_id" json:"role_id"`
    ResourceId int64 `form:"resource_id" json:"resource_id"`
    CreatedAt int64 `xorm:"created" form:"created_at" json:"created_at"`
    UpdatedAt int64 `xorm:"updated" form:"updated_at" json:"updated_at"`
    CreatedCustomerId  int64 `form:"created_customer_id" json:"created_customer_id"`
}

func (roleResource RoleResource) TableName() string {
    return "role_resource"
}
// 查询role 对应的所有的resource信息，以及选中的信息
func RoleResourceAllAndSelect(c *gin.Context){
    // 从resourceGroups，得到id 和 name对应的map - rGs
    rGs := make(MapInt64Str)
    resourceGroups, err := GetResourceGroupAll()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    for i:=0; i<len(resourceGroups); i++ {
        resourceGroup := resourceGroups[i];
        rGs[resourceGroup.Id] = resourceGroup.Name
    }
    
    // 得到选中的role
    // var resourceChecked []int64
    roleResources, err := GetCurrentRoleResourceAll(c)
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    vueMutilSelect := make(VueMutilSelect)
    // 得到所有的resource列表
    // var allResource []MapStrInterface
    resources, err := GetResourceAll()
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return  
    }
    for i:=0; i<len(resources); i++ {
        resource := resources[i];
        rGId := resource.GroupId
        rGName := rGs[rGId]
        if rGName != "" {
            checked := false
            for j:=0; j<len(roleResources); j++ {
                roleResource := roleResources[j];
                if roleResource.ResourceId == resource.Id {
                    checked = true
                }
            }
            vueMutilSelect[rGName] = append(vueMutilSelect[rGName], MapStrInterface{
                "id": resource.Id,
                "name": resource.Name,
                "checked": checked,
            })
        }
    }
    result := util.BuildSuccessResult(gin.H{
        "allResource": vueMutilSelect,
    })
    c.JSON(http.StatusOK, result)
}

// 通过 own_id 和role_ids 得到 resource_ids
func GetResourceIdsByRoles(role_ids []int64) ([]int64, error){
    var roleResources []RoleResource
    var resource_ids []int64
    err := engine.In("role_id", role_ids).Find(&roleResources) 
    if err != nil{
        return resource_ids, err 
    }
    for i:=0; i<len(roleResources); i++ {
        roleResource := roleResources[i]
        resource_ids = append(resource_ids, roleResource.ResourceId)
    }
    // resource_ids = helper.SliceInt64Unique(resource_ids)
    return resource_ids, nil 
}

// 根据role_id own_id 得到权限资源list
func GetCurrentRoleResourceAll(c *gin.Context) ([]RoleResource, error){
    var roleResources []RoleResource
    
    // role_id, err := strconv.Atoi(c.Param("role_id"))
    role_id, err := strconv.Atoi(c.DefaultQuery("role_id", ""))
    if err != nil{
        return  roleResources, err
    }
    log.Println( role_id)
    err = engine.Where("role_id = ? ", role_id).Find(&roleResources) 
    if err != nil{
        return roleResources, err 
    }
    log.Println( roleResources)
    return roleResources, nil
}

// 接收更新role resource的类型
type UpdateResource struct{
    RoleId int64 `form:"role_id" json:"role_id" binding:"required"`
    Resources []int64 `form:"resources" json:"resources" binding:"required"`
}
// 更新role 对应的resource
func RoleResourceUpdate(c *gin.Context){
    var updateResource UpdateResource
    err := c.ShouldBindJSON(&updateResource);
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    role_id := updateResource.RoleId
    resource_ids := updateResource.Resources
    if err != nil{
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 删除 在RoleResource表中role_id 和 own_id对应的所有的资源
    var roleResource RoleResource
    _, err = engine.Where("role_id = ?  ", role_id).Delete(&roleResource)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(err.Error()))
        return
    }
    // 将获取的数据插入。
    createdCustomerId := helper.GetCurrentCustomerId(c)
    for i:=0; i<len(resource_ids); i++ {
        var rr RoleResource
        rr.CreatedCustomerId = createdCustomerId
        rr.RoleId = role_id
        rr.ResourceId = resource_ids[i]
        _, err = engine.Insert(&rr)
    }
    result := util.BuildSuccessResult(gin.H{
        "updateResource":updateResource,
    })
    c.JSON(http.StatusOK, result)
}
