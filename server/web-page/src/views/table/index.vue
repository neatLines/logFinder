<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="20px">
      <el-form-item>
        <el-input v-model="form.filter" placeholder="过滤规则，like '2019*'，如果有多条规则使用\bb分割"/>
      </el-form-item>
      <el-form-item>
        <el-select v-model="form.appName" placeholder="选择应用">
          <el-option
            v-for="item in appName"
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
            v-for="item in applications[form.appName]"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
        <el-date-picker v-model="form.startTime" type="datetime" value-format="yyyy-MM-dd HH:mm:ss" placeholder="选择开始日期时间"></el-date-picker>
        <el-date-picker v-model="form.endTime" :disabled="form.needFlush" type="datetime" value-format="yyyy-MM-dd HH:mm:ss" placeholder="选择截止日期时间"></el-date-picker>
        <el-switch
          v-model="form.needFlush"
          active-text="实时刷新"
          active-color="#13ce66"
          inactive-color="#ff4949"
        ></el-switch>
        <el-button type="primary" @click="onSubmit">过滤</el-button>
      </el-form-item>
    </el-form>
    <el-tabs>
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
          <template slot-scope="scope">
            <el-popover trigger="hover" placement="top">
              <p>message: {{ scope.row.message }}</p>
              <div slot="reference" class="name-wrapper">{{ scope.row.message }}</div>
            </el-popover>
          </template>
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
        startTime: null,
        endTime: null,
        filter: "",
        needFlush: false,
        appName: null
      },
      appName: [],
      applications: {}
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
      this.list = [];
      this.websocketsend(JSON.stringify(this.form));
    },
    initWebSocket() {
      //初始化weosocket
      this.$store
        .dispatch("GetHosts")
        .then(response => {
          for (const key in response) {
            if (response.hasOwnProperty(key)) {
              let tmpvalue = [];
              for (const hostname in response[key]["hostname"]) {
                if (response[key]["hostname"].hasOwnProperty(hostname)) {
                  tmpvalue.push({ label: hostname, value: hostname });
                }
              }
              this.applications[key]=tmpvalue;
              this.appName.push({label: key, value: key})
            }
          }
          this.loading = false;
        })
        .catch(e => {
          this.loading = false;
        });
      this.listLoading = true;
      // const wsuri = "ws://localhost:8080/v1/ws";
      const wsuri = "ws://"+window.location.host+"/v1/ws";
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
