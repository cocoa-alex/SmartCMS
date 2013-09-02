        <div class="panel panel-default">
          <div class="panel-body">
          <a href="/">管理页面</a><span class="span-splite">/</span>用户模块
          </div>
        </div>
      <div class="bs-table">
        <table class="table table-hover">
        {{if .items}}
        <thead>
          <tr>
            <th>#</th>
            <th>用户名称</th>
            <th>用户密码</th>
            <th>注册时间</th>
          </tr>
        </thead>
        <tbody>
          {{range .items}}
          <tr>
            <td>{{ .Uid}}</td>
            <td>{{ .Username}}</td>
            <td>{{ .Password}}</td>
            <td>{{date2string .Created}}</td>
          </tr>
          {{end}}
        </tbody>
        </table>
        {{else}}
        <div class="alert alert-block ">
              <button type="button" class="close" data-dismiss="alert">×</button>
              <h4 class="alert-heading">Sorry!</h4>
              <p>亲，暂时没有数据啊，点击下面的按钮添加哦！！</p>
            </div>
                  <a class="btn btn-success" href="/admin/usr/add">
                    <i class="icon-pencil icon-white"></i>新增加一条
                  </a>            
          </div>
        {{end}}
      </div>