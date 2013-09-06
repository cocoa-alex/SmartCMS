        <div class="panel panel-default">
          <div class="panel-body">
            管理页面<span class="span-splite">/</span>用户模块
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
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          {{range .items}}
          <tr>
            <td>{{ .Uid}}</td>
            <td>{{ .Username}}</td>
            <td>{{ .Password}}</td>
            <td>{{date2string .Created}}</td>
            <td>
            <a data-toggle="modal" href="#myModal" class="btn"><span class="glyphicon glyphicon-remove"></span></a>
            <a href="#" class="btn"><span class="glyphicon glyphicon-plus"></span></a>
            <a href="#" class="btn"><span class="glyphicon glyphicon-pencil"></span></a>
            </td>
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
        <!-- Modal -->
        <div class="modal fade" id="myModal" role="dialog">
          <div class="modal-dialog">
            <div class="modal-content">
              <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title">{{ .usrttitle}}</h4>
              </div>
              <div class="modal-body">
                确定删除{{ .usrttitle}}？
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button type="button" class="btn btn-primary">确认删除</button>
              </div>
            </div><!-- /.modal-content -->
          </div><!-- /.modal-dialog -->
        </div><!-- /.modal -->