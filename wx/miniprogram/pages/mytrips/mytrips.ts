import { rental } from "../../service/proto_gen/rental/rental_pb"
import { TripService } from "../../service/proto_gen/trip"
import { formatDuration, formatFee } from "../../utils/format"
import { routing } from "../../utils/routing"
import { avatarUrlKey } from "../../utils/wxapi"
interface Trip{
    id:string
    shortId:string
    start:string
    end:string
    duration:string
    fee:string
    distance:string
    status:string
}
interface MainItem{
    id:string
    navId:string
    navScrollId:string
    data:Trip
}
interface NavItem{
    id:string
    mainId:string
    label:string
}
interface MainItemQueryResult{
    id:string
    top:number
    dataset:{
        navId:string
        navScrollId:string
    }
}
const tripStatusMap = new Map([
    [rental.v1.TripStatus.IN_PROGRESS, '进行中'],
    [rental.v1.TripStatus.FINISHED, '已完成'],
])
// pages/mytrips/mytrips.ts
Page({

  /**
   * 页面的初始数据
   */
  scrollStates:{
    mainItems:[] as MainItemQueryResult[]
},
  data: {
    avatarURL:"",
    
    promotionItems: [
        {
            img: 'https://img.mukewang.com/5f7301d80001fdee18720764.jpg',
            promotionID: 1,
        },            
        {
            img: 'https://img.mukewang.com/5f6805710001326c18720764.jpg',
            promotionID: 2,
        },
        {
            img: 'https://img.mukewang.com/5f6173b400013d4718720764.jpg',
            promotionID: 3,
        },
        {
            img: 'https://img.mukewang.com/5f7141ad0001b36418720764.jpg',
            promotionID: 4,
        },
    ],
    mainItems:[] as MainItem[],
    navScroll:'',
    navItems:[] as NavItem[],
    mainScroll:'',
    navCount:0,
    navSel:'',
  },
  onRegisterTap(){
    wx.navigateTo({
      url:routing.register(),
    })
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad() {
      
    this.populateTrips()
    this.prepareScrollStates()
    this.setData({
        avatarURL:wx.getStorageSync(avatarUrlKey)
    })
  },
  populateTrips(){
    const mainItems:MainItem[]=[]
    const navItems:NavItem[]=[]
    let navSel=''
    let prevNav=''
    TripService.GetTrips().then(trips=>{
        let i=0
        let mainId=''
        let navId=''
        let shortId=''
        trips.trips!.forEach(trip => {
            mainId='main-'+i
            navId='nav-'+i
            shortId=trip.id?.substr(trip.id.length-6)!
            if(!prevNav){
                prevNav=navId
            }
            const tripData:Trip={
                id:trip.id!,
                shortId:'****'+shortId,
                start:trip.trip?.start?.poiName||'未知',
                end:'',
                distance:'',
                duration:'',
                fee:'',
                status:tripStatusMap.get(trip.trip?.status!)||'未知',
            }
            const end=trip.trip?.end
            if(end){
                tripData.end=end.poiName||'未知'
                tripData.distance=end.kmDriven?.toFixed(1)+'公里'
                const dur=formatDuration((end.timestampSec||0) - (trip.trip?.start?.timestampSec||0))
                tripData.duration=`${dur.hh}时${dur.mm}分`
                tripData.fee=formatFee(end.feeCent||0)
            }
            mainItems.push({
                id:mainId,
                navId:navId,
                navScrollId:prevNav,
                data:tripData,
            })
            navItems.push({
                id:navId,
                mainId:mainId,
                label:shortId||'',
            })
            if(i===0){
                navSel=navId
            }
            prevNav=navId
            i++
        });
        this.setData({
            mainItems,
            navItems,
            navSel,
        },()=>{
            this.prepareScrollStates()
        })
    })
       
  },
  prepareScrollStates (){
    wx.createSelectorQuery().selectAll('.main-item').fields({
        id:true,
        dataset:true,
        rect:true,
    }).exec(res=>{
        this.scrollStates.mainItems=res[0]
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady() {
    wx.createSelectorQuery().select('#heading').boundingClientRect(rect=>{
        const height=wx.getSystemInfoSync().windowHeight-rect.height
        this.setData({
            tripsHeight:height,
            navCount:Math.round(height/50),
        })
    }).exec()
  },
  onPromotionItemTap(e: any) {
    const promotionID:number = e.currentTarget.dataset.promotionId
    if (promotionID) {
        console.log('claiming promotion', promotionID)
    }
    },
    onNavItemTap(e:any){
        const mainId:string=e.currentTarget?.dataset?.mainId
        const navId:string=e.currentTarget?.id
        if(mainId){
            this.setData({
                mainScroll:mainId,
                navSel:navId,
            })
        }

    },
    onMainScroll(e:any){
        const top:number=e.currentTarget.offsetTop+e.detail.scrollTop
        if(top===undefined){
            return
        }
        const selItem= this.scrollStates.mainItems.find((v => v.top>=top))
        if(!selItem){
            return
        }
        this.setData({
            navSel:selItem.dataset.navId,
            navScroll:selItem.dataset.navScrollId,
        })
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