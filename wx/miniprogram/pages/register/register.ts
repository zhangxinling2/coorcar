// pages/register/register.ts
Page({

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
    status:'UnSubmitted' as 'UnSubmitted'|'Pending'|'SubmittedFailed'|'Verified',
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
    this.setData({
      status:'Pending',
    })
    setTimeout(()=>{
      this.Verified()
    },3000)
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
    wx.redirectTo({
      url:'/pages/lock/lock',
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {

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