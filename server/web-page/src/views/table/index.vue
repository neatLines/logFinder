<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="20px">
      <el-form-item>
        <el-input v-model="form.filter" placeholder="过滤规则，like '2019*'，如果有多条规则使用\bb分割"/>
      </el-form-item>
      <el-form-item>
        <el-select v-model="appName" placeholder="选择应用">
          <el-option
            v-for="item in applications"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-select
          v-model="form.hostname"
          placeholder="选择主机名"
          multiple
          filterable
          allow-create
          default-first-option
        >
          <el-option
            v-for="item in appName"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-date-picker
          v-model="form.limitTime"
          type="datetimerange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        ></el-date-picker>
        <el-button type="primary" @click="onSubmit">过滤</el-button>
      </el-form-item>
    </el-form>
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-table
        v-loading="listLoading"
        :data="list"
        element-loading-text="Loading"
        border
        fit
        highlight-current-row
      >
        <el-table-column align="center" label="ID" width="95">
          <template slot-scope="scope">{{ scope.$index }}</template>
        </el-table-column>
        <el-table-column label="message">
          <template slot-scope="scope">{{ scope.row.message }}</template>
        </el-table-column>
        <el-table-column align="center" prop="created_at" label="time" width="200">
          <template slot-scope="scope">
            <i class="el-icon-time"/>
            <span>{{ scope.row.time }}</span>
          </template>
        </el-table-column>
      </el-table>
    </el-tabs>
  </div>
</template>

<script>

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: "success",
        draft: "gray",
        deleted: "danger"
      };
      return statusMap[status];
    }
  },
  data() {
    return {
      list: [],
      listLoading: true,
      form: {
        hostname: null,
        limitTime: [
          new Date(new Date().setDate(new Date().getDate() - 1)),
          new Date()
        ],
        filter: ""
      },
      appName: [],
      applications: [
        {
          value: [
            {
              value: "Beijing",
              label: "北京"
            },
            {
              value: "Shanghai",
              label: "上海"
            },
            {
              value: "Nanjing",
              label: "南京"
            },
            {
              value: "Chengdu",
              label: "成都"
            },
            {
              value: "Shenzhen",
              label: "深圳"
            },
            {
              value: "Guangzhou",
              label: "广州"
            }
          ],
          label: "黄金糕"
        },
        {
          value: [],
          label: "双皮奶"
        },
        {
          value: [],
          label: "蚵仔煎"
        },
        {
          value: [],
          label: "龙须面"
        },
        {
          value: [],
          label: "北京烤鸭"
        }
      ]
    };
  },
  created() {
    this.initWebSocket();
  },
  destroyed() {
    this.websock.close(); //离开路由之后断开websocket连接
  },
  methods: {
    onSubmit() {
      this.axios.get("/v1/hosts").then(response => {
        console.log(response.data);
      });
      console.log(this.form.hostname);
      this.list = [];
      this.websocketsend(JSON.stringify(this.form));
    },
    initWebSocket() {
      //初始化weosocket
      this.listLoading = true;
      const wsuri = "ws://localhost:8080/v1/ws";
      // const wsuri = "ws://"+window.location.host+"/v1/ws";
      this.websock = new WebSocket(wsuri);
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onopen = this.websocketonopen;
      this.websock.onerror = this.websocketonerror;
      this.websock.onclose = this.websocketclose;
    },
    websocketonopen() {
      //连接建立之后执行send方法发送数据
      this.listLoading = false;
    },
    websocketonerror() {
      //连接建立失败重连
      this.initWebSocket();
    },
    websocketonmessage(e) {
      //数据接收
      console.log(JSON.parse(e.data));
      this.list.push(JSON.parse(e.data));
    },
    websocketsend(Data) {
      //数据发送
      this.websock.send(Data);
    },
    websocketclose(e) {
      //关闭
      console.log("断开连接", e);
    }
  }
};
</script>
