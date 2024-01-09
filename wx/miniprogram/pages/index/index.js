"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const routing_1 = require("../../utils/routing");
const app = getApp();
const defaultAvatarUrl = 'https://mmbiz.qpic.cn/mmbiz/icTdbqWNOwNRna42FI242Lcia07jQodd2FJGIYQfG0LAJGFxM4FbnQP6yfMxBgJ0F3YRqJCJ1aPAK2dQagdusBZg/0';
Page({
    isPageShow: false,
    data: {
        setting: {
            skew: 0,
            rotate: 0,
            showLocation: true,
            showScale: true,
            subKey: '',
            layerStyle: -1,
            enableZoon: true,
            enableScroll: true,
            enableRotate: false,
            showCompass: false,
            enable3D: false,
            enableOverlooking: false,
            enableSatellite: false,
            enableTraffic: false,
        },
        location: {
            latitude: 31,
            longitude: 120,
        },
        scal: 10,
        markers: [{
                iconPath: "/resources/car.png",
                id: 0,
                latitude: 23.09994,
                longitude: 113.324520,
                width: 50,
                height: 50,
            }, {
                iconPath: "/resources/car.png",
                id: 1,
                latitude: 23.09994,
                longitude: 114.324520,
                width: 50,
                height: 50,
            }],
    },
    onHide() {
        this.isPageShow = false;
    },
    onShow() {
        this.isPageShow = true;
    },
    onMyLocationTap() {
        wx.getLocation({
            type: 'gcj02',
            success: res => {
                this.setData({
                    location: {
                        latitude: res.latitude,
                        longitude: res.longitude,
                    },
                });
            },
            fail: () => {
                wx.showToast({
                    icon: 'none',
                    title: '请前往设置页设置',
                });
            }
        });
    },
    moveCars() {
        const map = wx.createMapContext("map");
        const dest = {
            lagtitude: 23.09994,
            longitude: 113.324520,
        };
        const moveCar = () => {
            dest.lagtitude += 0.01,
                dest.longitude += 0.01,
                map.translateMarker({
                    autoRotate: false,
                    destination: {
                        latitude: dest.lagtitude,
                        longitude: dest.longitude,
                    },
                    markerId: 0,
                    rotate: 0,
                    duration: 5000,
                    animationEnd: () => {
                        if (this.isPageShow) {
                            moveCar();
                        }
                    }
                });
        };
        moveCar();
    },
    onScanTap() {
        wx.scanCode({
            success: () => {
                const carId = 'car123';
                const redirectURL = routing_1.routing.lock({
                    car_id: carId
                });
                wx.navigateTo({
                    url: routing_1.routing.register({
                        redirectURL: redirectURL
                    })
                });
            },
            fail: console.error
        });
    },
    onMyTripsTap() {
        wx.navigateTo({
            url: routing_1.routing.mytrips(),
        });
    }
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiaW5kZXguanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJpbmRleC50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiOztBQUdBLGlEQUE2QztBQUc3QyxNQUFNLEdBQUcsR0FBRyxNQUFNLEVBQWMsQ0FBQTtBQUNoQyxNQUFNLGdCQUFnQixHQUFHLHdIQUF3SCxDQUFBO0FBRWpKLElBQUksQ0FBQztJQUNILFVBQVUsRUFBQyxLQUFLO0lBQ2hCLElBQUksRUFBQztRQUNILE9BQU8sRUFBQztZQUNKLElBQUksRUFBQyxDQUFDO1lBQ04sTUFBTSxFQUFDLENBQUM7WUFDUixZQUFZLEVBQUMsSUFBSTtZQUNqQixTQUFTLEVBQUMsSUFBSTtZQUNkLE1BQU0sRUFBQyxFQUFFO1lBQ1QsVUFBVSxFQUFDLENBQUMsQ0FBQztZQUNiLFVBQVUsRUFBQyxJQUFJO1lBQ2YsWUFBWSxFQUFDLElBQUk7WUFDakIsWUFBWSxFQUFDLEtBQUs7WUFDbEIsV0FBVyxFQUFDLEtBQUs7WUFDakIsUUFBUSxFQUFDLEtBQUs7WUFDZCxpQkFBaUIsRUFBQyxLQUFLO1lBQ3ZCLGVBQWUsRUFBQyxLQUFLO1lBQ3JCLGFBQWEsRUFBQyxLQUFLO1NBQ3RCO1FBQ0QsUUFBUSxFQUFDO1lBQ0wsUUFBUSxFQUFDLEVBQUU7WUFDWCxTQUFTLEVBQUMsR0FBRztTQUNoQjtRQUNELElBQUksRUFBQyxFQUFFO1FBQ1AsT0FBTyxFQUFDLENBQUM7Z0JBQ0wsUUFBUSxFQUFDLG9CQUFvQjtnQkFDN0IsRUFBRSxFQUFDLENBQUM7Z0JBQ0osUUFBUSxFQUFDLFFBQVE7Z0JBQ2pCLFNBQVMsRUFBQyxVQUFVO2dCQUNwQixLQUFLLEVBQUMsRUFBRTtnQkFDUixNQUFNLEVBQUMsRUFBRTthQUNaLEVBQUM7Z0JBQ0UsUUFBUSxFQUFDLG9CQUFvQjtnQkFDN0IsRUFBRSxFQUFDLENBQUM7Z0JBQ0osUUFBUSxFQUFDLFFBQVE7Z0JBQ2pCLFNBQVMsRUFBQyxVQUFVO2dCQUNwQixLQUFLLEVBQUMsRUFBRTtnQkFDUixNQUFNLEVBQUMsRUFBRTthQUNaLENBQUM7S0FDSDtJQUNELE1BQU07UUFDSixJQUFJLENBQUMsVUFBVSxHQUFDLEtBQUssQ0FBQTtJQUN2QixDQUFDO0lBQ0QsTUFBTTtRQUNKLElBQUksQ0FBQyxVQUFVLEdBQUMsSUFBSSxDQUFBO0lBQ3RCLENBQUM7SUFDRCxlQUFlO1FBQ2IsRUFBRSxDQUFDLFdBQVcsQ0FBQztZQUNiLElBQUksRUFBQyxPQUFPO1lBQ1osT0FBTyxFQUFDLEdBQUcsQ0FBQSxFQUFFO2dCQUNYLElBQUksQ0FBQyxPQUFPLENBQUM7b0JBQ1gsUUFBUSxFQUFDO3dCQUNQLFFBQVEsRUFBQyxHQUFHLENBQUMsUUFBUTt3QkFDckIsU0FBUyxFQUFDLEdBQUcsQ0FBQyxTQUFTO3FCQUN4QjtpQkFDRixDQUFDLENBQUE7WUFDSixDQUFDO1lBQ0QsSUFBSSxFQUFDLEdBQUUsRUFBRTtnQkFDUCxFQUFFLENBQUMsU0FBUyxDQUFDO29CQUNYLElBQUksRUFBQyxNQUFNO29CQUNYLEtBQUssRUFBQyxVQUFVO2lCQUNqQixDQUFDLENBQUE7WUFDSixDQUFDO1NBQ0YsQ0FBQyxDQUFBO0lBQ0osQ0FBQztJQUNELFFBQVE7UUFDTixNQUFNLEdBQUcsR0FBRSxFQUFFLENBQUMsZ0JBQWdCLENBQUMsS0FBSyxDQUFDLENBQUE7UUFDckMsTUFBTSxJQUFJLEdBQUM7WUFDVCxTQUFTLEVBQUMsUUFBUTtZQUNsQixTQUFTLEVBQUMsVUFBVTtTQUNyQixDQUFBO1FBQ0QsTUFBTSxPQUFPLEdBQUcsR0FBRSxFQUFFO1lBQ2xCLElBQUksQ0FBQyxTQUFTLElBQUUsSUFBSTtnQkFDcEIsSUFBSSxDQUFDLFNBQVMsSUFBRSxJQUFJO2dCQUNwQixHQUFHLENBQUMsZUFBZSxDQUFDO29CQUNsQixVQUFVLEVBQUMsS0FBSztvQkFDaEIsV0FBVyxFQUFDO3dCQUNWLFFBQVEsRUFBQyxJQUFJLENBQUMsU0FBUzt3QkFDdkIsU0FBUyxFQUFDLElBQUksQ0FBQyxTQUFTO3FCQUN6QjtvQkFDRCxRQUFRLEVBQUMsQ0FBQztvQkFDVixNQUFNLEVBQUMsQ0FBQztvQkFDUixRQUFRLEVBQUMsSUFBSTtvQkFDYixZQUFZLEVBQUMsR0FBRSxFQUFFO3dCQUNmLElBQUcsSUFBSSxDQUFDLFVBQVUsRUFBQzs0QkFDakIsT0FBTyxFQUFFLENBQUE7eUJBQ1Y7b0JBQ0gsQ0FBQztpQkFDRixDQUFDLENBQUE7UUFDSixDQUFDLENBQUE7UUFDRCxPQUFPLEVBQUUsQ0FBQTtJQUNYLENBQUM7SUFDRCxTQUFTO1FBQ1AsRUFBRSxDQUFDLFFBQVEsQ0FBQztZQUNWLE9BQU8sRUFBQyxHQUFFLEVBQUU7Z0JBQ1YsTUFBTSxLQUFLLEdBQUMsUUFBUSxDQUFBO2dCQUNwQixNQUFNLFdBQVcsR0FBQyxpQkFBTyxDQUFDLElBQUksQ0FBQztvQkFDN0IsTUFBTSxFQUFDLEtBQUs7aUJBQ2IsQ0FBQyxDQUFBO2dCQUNGLEVBQUUsQ0FBQyxVQUFVLENBQUM7b0JBQ1osR0FBRyxFQUFDLGlCQUFPLENBQUMsUUFBUSxDQUFDO3dCQUNuQixXQUFXLEVBQUMsV0FBVztxQkFDeEIsQ0FBQztpQkFDSCxDQUFDLENBQUE7WUFDSixDQUFDO1lBQ0QsSUFBSSxFQUFDLE9BQU8sQ0FBQyxLQUFLO1NBQ25CLENBQUMsQ0FBQTtJQUNKLENBQUM7SUFDRCxZQUFZO1FBQ1YsRUFBRSxDQUFDLFVBQVUsQ0FBQztZQUNaLEdBQUcsRUFBQyxpQkFBTyxDQUFDLE9BQU8sRUFBRTtTQUN0QixDQUFDLENBQUE7SUFDSixDQUFDO0NBQ0YsQ0FBQyxDQUFBIiwic291cmNlc0NvbnRlbnQiOlsiLy8gaW5kZXgudHNcclxuXHJcbmltcG9ydCB7IElBcHBPcHRpb24gfSBmcm9tIFwiLi4vLi4vYXBwb3B0aW9uXCJcclxuaW1wb3J0IHsgcm91dGluZyB9IGZyb20gXCIuLi8uLi91dGlscy9yb3V0aW5nXCJcclxuXHJcbi8vIOiOt+WPluW6lOeUqOWunuS+i1xyXG5jb25zdCBhcHAgPSBnZXRBcHA8SUFwcE9wdGlvbj4oKVxyXG5jb25zdCBkZWZhdWx0QXZhdGFyVXJsID0gJ2h0dHBzOi8vbW1iaXoucXBpYy5jbi9tbWJpei9pY1RkYnFXTk93TlJuYTQyRkkyNDJMY2lhMDdqUW9kZDJGSkdJWVFmRzBMQUpHRnhNNEZiblFQNnlmTXhCZ0owRjNZUnFKQ0oxYVBBSzJkUWFnZHVzQlpnLzAnXHJcblxyXG5QYWdlKHtcclxuICBpc1BhZ2VTaG93OmZhbHNlLFxyXG4gIGRhdGE6e1xyXG4gICAgc2V0dGluZzp7XHJcbiAgICAgICAgc2tldzowLFxyXG4gICAgICAgIHJvdGF0ZTowLFxyXG4gICAgICAgIHNob3dMb2NhdGlvbjp0cnVlLFxyXG4gICAgICAgIHNob3dTY2FsZTp0cnVlLFxyXG4gICAgICAgIHN1YktleTonJyxcclxuICAgICAgICBsYXllclN0eWxlOi0xLFxyXG4gICAgICAgIGVuYWJsZVpvb246dHJ1ZSxcclxuICAgICAgICBlbmFibGVTY3JvbGw6dHJ1ZSxcclxuICAgICAgICBlbmFibGVSb3RhdGU6ZmFsc2UsXHJcbiAgICAgICAgc2hvd0NvbXBhc3M6ZmFsc2UsXHJcbiAgICAgICAgZW5hYmxlM0Q6ZmFsc2UsXHJcbiAgICAgICAgZW5hYmxlT3Zlcmxvb2tpbmc6ZmFsc2UsXHJcbiAgICAgICAgZW5hYmxlU2F0ZWxsaXRlOmZhbHNlLFxyXG4gICAgICAgIGVuYWJsZVRyYWZmaWM6ZmFsc2UsXHJcbiAgICB9LFxyXG4gICAgbG9jYXRpb246e1xyXG4gICAgICAgIGxhdGl0dWRlOjMxLFxyXG4gICAgICAgIGxvbmdpdHVkZToxMjAsXHJcbiAgICB9LFxyXG4gICAgc2NhbDoxMCxcclxuICAgIG1hcmtlcnM6W3tcclxuICAgICAgICBpY29uUGF0aDpcIi9yZXNvdXJjZXMvY2FyLnBuZ1wiLFxyXG4gICAgICAgIGlkOjAsXHJcbiAgICAgICAgbGF0aXR1ZGU6MjMuMDk5OTQsXHJcbiAgICAgICAgbG9uZ2l0dWRlOjExMy4zMjQ1MjAsXHJcbiAgICAgICAgd2lkdGg6NTAsXHJcbiAgICAgICAgaGVpZ2h0OjUwLFxyXG4gICAgfSx7XHJcbiAgICAgICAgaWNvblBhdGg6XCIvcmVzb3VyY2VzL2Nhci5wbmdcIixcclxuICAgICAgICBpZDoxLFxyXG4gICAgICAgIGxhdGl0dWRlOjIzLjA5OTk0LFxyXG4gICAgICAgIGxvbmdpdHVkZToxMTQuMzI0NTIwLFxyXG4gICAgICAgIHdpZHRoOjUwLFxyXG4gICAgICAgIGhlaWdodDo1MCxcclxuICAgIH1dLFxyXG4gIH0sXHJcbiAgb25IaWRlKCl7XHJcbiAgICB0aGlzLmlzUGFnZVNob3c9ZmFsc2VcclxuICB9LFxyXG4gIG9uU2hvdygpIHtcclxuICAgIHRoaXMuaXNQYWdlU2hvdz10cnVlXHJcbiAgfSxcclxuICBvbk15TG9jYXRpb25UYXAoKXtcclxuICAgIHd4LmdldExvY2F0aW9uKHtcclxuICAgICAgdHlwZTonZ2NqMDInLFxyXG4gICAgICBzdWNjZXNzOnJlcz0+e1xyXG4gICAgICAgIHRoaXMuc2V0RGF0YSh7XHJcbiAgICAgICAgICBsb2NhdGlvbjp7XHJcbiAgICAgICAgICAgIGxhdGl0dWRlOnJlcy5sYXRpdHVkZSxcclxuICAgICAgICAgICAgbG9uZ2l0dWRlOnJlcy5sb25naXR1ZGUsXHJcbiAgICAgICAgICB9LFxyXG4gICAgICAgIH0pXHJcbiAgICAgIH0sXHJcbiAgICAgIGZhaWw6KCk9PntcclxuICAgICAgICB3eC5zaG93VG9hc3Qoe1xyXG4gICAgICAgICAgaWNvbjonbm9uZScsXHJcbiAgICAgICAgICB0aXRsZTon6K+35YmN5b6A6K6+572u6aG16K6+572uJyxcclxuICAgICAgICB9KVxyXG4gICAgICB9XHJcbiAgICB9KVxyXG4gIH0sXHJcbiAgbW92ZUNhcnMoKXtcclxuICAgIGNvbnN0IG1hcD0gd3guY3JlYXRlTWFwQ29udGV4dChcIm1hcFwiKVxyXG4gICAgY29uc3QgZGVzdD17XHJcbiAgICAgIGxhZ3RpdHVkZToyMy4wOTk5NCxcclxuICAgICAgbG9uZ2l0dWRlOjExMy4zMjQ1MjAsXHJcbiAgICB9XHJcbiAgICBjb25zdCBtb3ZlQ2FyID0gKCk9PntcclxuICAgICAgZGVzdC5sYWd0aXR1ZGUrPTAuMDEsXHJcbiAgICAgIGRlc3QubG9uZ2l0dWRlKz0wLjAxLFxyXG4gICAgICBtYXAudHJhbnNsYXRlTWFya2VyKHtcclxuICAgICAgICBhdXRvUm90YXRlOmZhbHNlLFxyXG4gICAgICAgIGRlc3RpbmF0aW9uOntcclxuICAgICAgICAgIGxhdGl0dWRlOmRlc3QubGFndGl0dWRlLFxyXG4gICAgICAgICAgbG9uZ2l0dWRlOmRlc3QubG9uZ2l0dWRlLFxyXG4gICAgICAgIH0sXHJcbiAgICAgICAgbWFya2VySWQ6MCxcclxuICAgICAgICByb3RhdGU6MCxcclxuICAgICAgICBkdXJhdGlvbjo1MDAwLFxyXG4gICAgICAgIGFuaW1hdGlvbkVuZDooKT0+e1xyXG4gICAgICAgICAgaWYodGhpcy5pc1BhZ2VTaG93KXtcclxuICAgICAgICAgICAgbW92ZUNhcigpXHJcbiAgICAgICAgICB9XHJcbiAgICAgICAgfVxyXG4gICAgICB9KVxyXG4gICAgfVxyXG4gICAgbW92ZUNhcigpXHJcbiAgfSxcclxuICBvblNjYW5UYXAoKXtcclxuICAgIHd4LnNjYW5Db2RlKHtcclxuICAgICAgc3VjY2VzczooKT0+e1xyXG4gICAgICAgIGNvbnN0IGNhcklkPSdjYXIxMjMnXHJcbiAgICAgICAgY29uc3QgcmVkaXJlY3RVUkw9cm91dGluZy5sb2NrKHtcclxuICAgICAgICAgIGNhcl9pZDpjYXJJZFxyXG4gICAgICAgIH0pXHJcbiAgICAgICAgd3gubmF2aWdhdGVUbyh7XHJcbiAgICAgICAgICB1cmw6cm91dGluZy5yZWdpc3Rlcih7XHJcbiAgICAgICAgICAgIHJlZGlyZWN0VVJMOnJlZGlyZWN0VVJMXHJcbiAgICAgICAgICB9KVxyXG4gICAgICAgIH0pXHJcbiAgICAgIH0sXHJcbiAgICAgIGZhaWw6Y29uc29sZS5lcnJvclxyXG4gICAgfSlcclxuICB9LFxyXG4gIG9uTXlUcmlwc1RhcCgpe1xyXG4gICAgd3gubmF2aWdhdGVUbyh7XHJcbiAgICAgIHVybDpyb3V0aW5nLm15dHJpcHMoKSxcclxuICAgIH0pXHJcbiAgfVxyXG59KVxyXG4iXX0=