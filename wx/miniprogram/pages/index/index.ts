// index.ts
// 获取应用实例
const app = getApp<IAppOption>()
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0'

Page({
  isPageShow:false,
  data:{
    setting:{
        skew:0,
        rotate:0,
        showLocation:true,
        showScale:true,
        subKey:'',
        layerStyle:-1,
        enableZoon:true,
        enableScroll:true,
        enableRotate:false,
        showCompass:false,
        enable3D:false,
        enableOverlooking:false,
        enableSatellite:false,
        enableTraffic:false,
    },
    location:{
        latitude:31,
        longitude:120,
    },
    scal:10,
    markers:[{
        iconPath:"/resources/car.png",
        id:0,
        latitude:23.09994,
        longitude:113.324520,
        width:50,
        height:50,
    },{
        iconPath:"/resources/car.png",
        id:1,
        latitude:23.09994,
        longitude:114.324520,
        width:50,
        height:50,
    }],
  },
  onHide(){
    this.isPageShow=false
  },
  onShow() {
    this.isPageShow=true
  },
  onMyLocationTap(){
    wx.getLocation({
      type:'gcj02',
      success:res=>{
        this.setData({
          location:{
            latitude:res.latitude,
            longitude:res.longitude,
          },
        })
      },
      fail:()=>{
        wx.showToast({
          icon:'none',
          title:'请前往设置页设置',
        })
      }
    })
  },
  moveCars(){
    const map= wx.createMapContext("map")
    const dest={
      lagtitude:23.09994,
      longitude:113.324520,
    }
    const moveCar = ()=>{
      dest.lagtitude+=0.01,
      dest.longitude+=0.01,
      map.translateMarker({
        autoRotate:false,
        destination:{
          latitude:dest.lagtitude,
          longitude:dest.longitude,
        },
        markerId:0,
        rotate:0,
        duration:5000,
        animationEnd:()=>{
          if(this.isPageShow){
            moveCar()
          }
        }
      })
    }
    moveCar()
  }
})
