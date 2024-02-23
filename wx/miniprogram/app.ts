import camelcaseKeys from "camelcase-keys"
import { IAppOption } from "./appoption"
import { auth } from "./service/proto_gen/auth_pb"
import { rental } from "./service/proto_gen/rental_pb"

// app.ts
App<IAppOption>({
  globalData: {
  },
  async onLaunch() {
      // wx.request({
      //     url:"http://localhost:8080/trip/trip123",
      //     method:"GET",
      //     success:res=>{
      //       const getTripResp=coolcar.GetTripResponse.fromObject(camelcaseKeys(res.data as object,{deep:true}))
      //       console.log(getTripResp)
      //       console.log(coolcar.TripStatus[getTripResp.trip?.status!])
      //     },
      //     fail:console.error,
      // })
    
    console.log("完成网络请求")
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: res => {
        wx.request({
          url:"http://localhost:8080/v1/auth/login",
          method:"POST",
          data:{code:res.code},
          success:res=>{
            const loginResp:auth.v1.ILoginResponse = auth.v1.LoginResponse.fromObject(camelcaseKeys(res.data as object,{deep:true}))
            wx.request({
                url:'http://localhost:8080/v1/trip/create',
                method:'POST',
                data:{
                    start:'abc'
                } as rental.v1.ICreateTripRequest,
                header:{
                    authorization:'Bearer '+loginResp.accessToken,
                }
            })
          }
        })
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      },
    })

   
  },
})