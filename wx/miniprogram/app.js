"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const trip_pb_d_1 = require("./service/proto_gen/trip_pb.d");
App({
    globalData: {},
    onLaunch() {
        wx.request({
            url: "http://localhost:8080/trip/trip123",
            method: "GET",
            success: res => {
                const getTripResp = trip_pb_d_1.coolcar.GetTripResponse.fromObject(res.data);
                console.log(getTripResp);
            },
            fail: console.error,
            complete: console.log,
        });
        console.log("完成网络请求");
        const logs = wx.getStorageSync('logs') || [];
        logs.unshift(Date.now());
        wx.setStorageSync('logs', logs);
        wx.login({
            success: res => {
                console.log(res.code);
            },
        });
    },
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYXBwLmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsiYXBwLnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7O0FBQ0EsNkRBQXVEO0FBR3ZELEdBQUcsQ0FBYTtJQUNkLFVBQVUsRUFBRSxFQUNYO0lBQ0QsUUFBUTtRQUNKLEVBQUUsQ0FBQyxPQUFPLENBQUM7WUFDUCxHQUFHLEVBQUMsb0NBQW9DO1lBQ3hDLE1BQU0sRUFBQyxLQUFLO1lBQ1osT0FBTyxFQUFDLEdBQUcsQ0FBQSxFQUFFO2dCQUNYLE1BQU0sV0FBVyxHQUFDLG1CQUFPLENBQUMsZUFBZSxDQUFDLFVBQVUsQ0FBQyxHQUFHLENBQUMsSUFBYyxDQUFDLENBQUE7Z0JBQ3hFLE9BQU8sQ0FBQyxHQUFHLENBQUMsV0FBVyxDQUFDLENBQUE7WUFDMUIsQ0FBQztZQUNELElBQUksRUFBQyxPQUFPLENBQUMsS0FBSztZQUNsQixRQUFRLEVBQUMsT0FBTyxDQUFDLEdBQUc7U0FDdkIsQ0FBQyxDQUFBO1FBQ0osT0FBTyxDQUFDLEdBQUcsQ0FBQyxRQUFRLENBQUMsQ0FBQTtRQUVyQixNQUFNLElBQUksR0FBRyxFQUFFLENBQUMsY0FBYyxDQUFDLE1BQU0sQ0FBQyxJQUFJLEVBQUUsQ0FBQTtRQUM1QyxJQUFJLENBQUMsT0FBTyxDQUFDLElBQUksQ0FBQyxHQUFHLEVBQUUsQ0FBQyxDQUFBO1FBQ3hCLEVBQUUsQ0FBQyxjQUFjLENBQUMsTUFBTSxFQUFFLElBQUksQ0FBQyxDQUFBO1FBRy9CLEVBQUUsQ0FBQyxLQUFLLENBQUM7WUFDUCxPQUFPLEVBQUUsR0FBRyxDQUFDLEVBQUU7Z0JBQ2IsT0FBTyxDQUFDLEdBQUcsQ0FBQyxHQUFHLENBQUMsSUFBSSxDQUFDLENBQUE7WUFFdkIsQ0FBQztTQUNGLENBQUMsQ0FBQTtJQUdKLENBQUM7Q0FDRixDQUFDLENBQUEiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBJQXBwT3B0aW9uIH0gZnJvbSBcIi4vYXBwb3B0aW9uXCJcclxuaW1wb3J0IHsgY29vbGNhciB9IGZyb20gXCIuL3NlcnZpY2UvcHJvdG9fZ2VuL3RyaXBfcGIuZFwiXHJcblxyXG4vLyBhcHAudHNcclxuQXBwPElBcHBPcHRpb24+KHtcclxuICBnbG9iYWxEYXRhOiB7XHJcbiAgfSxcclxuICBvbkxhdW5jaCgpIHtcclxuICAgICAgd3gucmVxdWVzdCh7XHJcbiAgICAgICAgICB1cmw6XCJodHRwOi8vbG9jYWxob3N0OjgwODAvdHJpcC90cmlwMTIzXCIsXHJcbiAgICAgICAgICBtZXRob2Q6XCJHRVRcIixcclxuICAgICAgICAgIHN1Y2Nlc3M6cmVzPT57XHJcbiAgICAgICAgICAgIGNvbnN0IGdldFRyaXBSZXNwPWNvb2xjYXIuR2V0VHJpcFJlc3BvbnNlLmZyb21PYmplY3QocmVzLmRhdGEgYXMgb2JqZWN0KVxyXG4gICAgICAgICAgICBjb25zb2xlLmxvZyhnZXRUcmlwUmVzcClcclxuICAgICAgICAgIH0sXHJcbiAgICAgICAgICBmYWlsOmNvbnNvbGUuZXJyb3IsXHJcbiAgICAgICAgICBjb21wbGV0ZTpjb25zb2xlLmxvZyxcclxuICAgICAgfSlcclxuICAgIGNvbnNvbGUubG9nKFwi5a6M5oiQ572R57uc6K+35rGCXCIpXHJcbiAgICAvLyDlsZXnpLrmnKzlnLDlrZjlgqjog73liptcclxuICAgIGNvbnN0IGxvZ3MgPSB3eC5nZXRTdG9yYWdlU3luYygnbG9ncycpIHx8IFtdXHJcbiAgICBsb2dzLnVuc2hpZnQoRGF0ZS5ub3coKSlcclxuICAgIHd4LnNldFN0b3JhZ2VTeW5jKCdsb2dzJywgbG9ncylcclxuXHJcbiAgICAvLyDnmbvlvZVcclxuICAgIHd4LmxvZ2luKHtcclxuICAgICAgc3VjY2VzczogcmVzID0+IHtcclxuICAgICAgICBjb25zb2xlLmxvZyhyZXMuY29kZSlcclxuICAgICAgICAvLyDlj5HpgIEgcmVzLmNvZGUg5Yiw5ZCO5Y+w5o2i5Y+WIG9wZW5JZCwgc2Vzc2lvbktleSwgdW5pb25JZFxyXG4gICAgICB9LFxyXG4gICAgfSlcclxuXHJcbiAgIFxyXG4gIH0sXHJcbn0pIl19