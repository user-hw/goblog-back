/*
 * @Author: HengweiXu 1761173100@qq.com
 * @Date: 2023-01-31 16:04:59
 * @LastEditors: HengweiXu 1761173100@qq.com
 * @LastEditTime: 2023-01-31 16:40:51
 * @FilePath: /goblog-back/dao/saveAndUpdatePost.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"fmt"
	model "goblog-end/model/DBmodel"
)

func SaveAndUpdatePost(blogPost model.BlogPost) (int, error) {

	// DB.Create(&blogPost)

	//一般项目中我们会类似下面的写法，通过Error对象检测，插入数据有没有成功，如果没有错误那就是数据写入成功了。
	if err := DB.Create(&blogPost).Error; err != nil {
		fmt.Println("插入失败", err)
		return -1, err
	} else {
		return 200, nil
	}

}
