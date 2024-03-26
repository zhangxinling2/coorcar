import { ProfileService } from "../../service/proto_gen/profile";
import { rental } from "../../service/proto_gen/rental/rental_pb";
import { routing } from "../../utils/routing";
function padString(n: number) {
    return n < 10 ? '0'+n.toFixed(0) : n.toFixed(0)
}
function formatDate(millis:number){
    const dt=new Date(millis)
    const y=dt.getFullYear()
    const m=dt.getMonth()+1
    const d=dt.getDay()
    return `${padString(y)}-${padString(m)}-${padString(d)}`
}
// pages/register/register.ts
Page({
  redirectURL:'',
  profileRefresher:0,
  /**
   * 页面的初始数据
   */
  data: {
    genderIndex:0,
    birthData:'2022-1-1',
    licNo:"",
    licName:"",
    genders:['未知','男','女','其它'],
    licImgURL:'',
    status:rental.v1.IdentityStatus[rental.v1.IdentityStatus.UNSUBMITTED],
  },
  rentalIdentity(i?:rental.v1.IIdentity){
    this.setData({
        licNo:i?.licNumber||'',
        licName:i?.name||'',
        genderIndex:i?.gender||0,
        birthData:formatDate(i?.birthDataMillis||0),
    })
  },
  rentalProfile(p:rental.v1.IProfile){
    this.rentalIdentity(p.identity!)
    this.setData({
        status:rental.v1.IdentityStatus[p.identityStatus||0],
    })
  },
  onUploadLic(){
    wx.chooseMedia({
      success:res=>{
        this.setData({
          licImgURL:res.tempFiles[0].tempFilePath
        })
      }
    })
    setTimeout(() => {
      this.setData({
        genderIndex:1,
        birthData:'2000-1-1',
        licNo:'200001010',
        licName:'张三',
      })
    }, 1000);
  },
  onGenderChange(e:any){
    this.setData({
      genderIndex:e.detail.value,
    })
  },
  onBirthChange(e:any){
    this.setData({
      birthData:e.detail.value,
    })
  },
  onSubmit(){
    ProfileService.SubmitProfile({
        licNumber:this.data.licNo,
        name:this.data.licName,
        gender:this.data.genderIndex,
        birthDataMillis:Date.parse(this.data.birthData),
    }).then(p=>{
        this.rentalProfile(p)
        this.scheduleProfileRefresher()
    })
  },
  scheduleProfileRefresher(){
    this.profileRefresher=setInterval(()=>{
        ProfileService.GetProfile().then(p=>{
            this.rentalProfile(p)
            if(p.identityStatus!==rental.v1.IdentityStatus.PENDING){
                this.clearProfileRefresher()
            }
            if(p.identityStatus===rental.v1.IdentityStatus.VERIFIED){
                this.onLicVerified()
            }
        })
    },1000)
  },
  clearProfileRefresher(){
    if(this.profileRefresher){
        clearInterval(this.profileRefresher)
        this.profileRefresher=0
    }
  },
  onLicVerified(){
    if(this.redirectURL){
        wx.redirectTo({
            url:this.redirectURL,
        })
    }
  },
  onRetry(){
    this.setData({
      status:'UnSubmitted',
      licImgURL:'',
    })
  },
  Verified(){
    // this.setData({
    //   states:'SubmittedFailed',
    // })
    this.setData({
      status:'Verified',
    })
    if(this.redirectURL){
      wx.redirectTo({
        url:this.redirectURL,
      })
    }
    
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad(opt:Record<'redirect',string>) {
    const o: routing.RegisterOpts =opt
    if(o.redirect){
      this.redirectURL=decodeURIComponent(o.redirect)
    }
    ProfileService.GetProfile().then(p=>{
       this.rentalIdentity(p.identity!)
    })
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide() {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload() {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage() {

  }
})