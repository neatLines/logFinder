<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="20px">
      <el-form-item>
        <el-input v-model="form.filter" placeholder="过滤规则，like '2019*'，如果有多条规则使用\bb分割"/>
      </el-form-item>
      <el-form-item>
        <el-select v-model="form.appName" placeholder="选择应用">
          <el-option
            v-for="item in form.applications"
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
            v-for="item in form.appName"
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
      <el-tab-pane label="Table" name="first">
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
          <el-table-column label="Title">
            <template slot-scope="scope">{{ scope.row.title }}</template>
          </el-table-column>
          <el-table-column label="Author" width="110" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.author }}</span>
            </template>
          </el-table-column>
          <el-table-column label="Pageviews" width="110" align="center">
            <template slot-scope="scope">{{ scope.row.pageviews }}</template>
          </el-table-column>
          <el-table-column class-name="status-col" label="Status" width="110" align="center">
            <template slot-scope="scope">
              <el-tag :type="scope.row.status | statusFilter">{{ scope.row.status }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="center" prop="created_at" label="Display_time" width="200">
            <template slot-scope="scope">
              <i class="el-icon-time"/>
              <span>{{ scope.row.display_time }}</span>
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
      <el-tab-pane label="Row" name="second"></el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
import { getList } from "@/api/table";

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
      list: null,
      listLoading: true,
              activeName: 'first',
      form: {
        hostname: null,
        appName: [],
        limitTime: [
          new Date(new Date().setDate(new Date().getDate() - 1)),
          new Date()
        ],
        filter: "",
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
      }
    };
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      getList(this.listQuery).then(response => {
        this.list = response.data.items;
        this.listLoading = false;
      });
    },
    onSubmit() {
      console.log(this.form.hostname);
    }
  }
};
</script>
