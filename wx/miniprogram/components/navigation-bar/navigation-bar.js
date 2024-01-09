"use strict";
Component({
    options: {
        multipleSlots: true
    },
    properties: {
        extClass: {
            type: String,
            value: ''
        },
        title: {
            type: String,
            value: ''
        },
        background: {
            type: String,
            value: ''
        },
        color: {
            type: String,
            value: ''
        },
        back: {
            type: Boolean,
            value: true
        },
        loading: {
            type: Boolean,
            value: false
        },
        homeButton: {
            type: Boolean,
            value: false,
        },
        animated: {
            type: Boolean,
            value: true
        },
        show: {
            type: Boolean,
            value: true,
            observer: '_showChange'
        },
        delta: {
            type: Number,
            value: 1
        },
    },
    data: {
        displayStyle: ''
    },
    lifetimes: {
        attached() {
            const rect = wx.getMenuButtonBoundingClientRect();
            wx.getSystemInfo({
                success: (res) => {
                    const isAndroid = res.platform === 'android';
                    const isDevtools = res.platform === 'devtools';
                    this.setData({
                        ios: !isAndroid,
                        innerPaddingRight: `padding-right: ${res.windowWidth - rect.left}px`,
                        leftWidth: `width: ${res.windowWidth - rect.left}px`,
                        safeAreaTop: isDevtools || isAndroid ? `height: calc(var(--height) + ${res.safeArea.top}px); padding-top: ${res.safeArea.top}px` : ``
                    });
                }
            });
        },
    },
    methods: {
        _showChange(show) {
            const animated = this.data.animated;
            let displayStyle = '';
            if (animated) {
                displayStyle = `opacity: ${show ? '1' : '0'};transition:opacity 0.5s;`;
            }
            else {
                displayStyle = `display: ${show ? '' : 'none'}`;
            }
            this.setData({
                displayStyle
            });
        },
        back() {
            const data = this.data;
            if (data.delta) {
                wx.navigateBack({
                    delta: data.delta
                });
            }
            this.triggerEvent('back', { delta: data.delta }, {});
        }
    },
});
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibmF2aWdhdGlvbi1iYXIuanMiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJuYXZpZ2F0aW9uLWJhci50cyJdLCJuYW1lcyI6W10sIm1hcHBpbmdzIjoiO0FBQUEsU0FBUyxDQUFDO0lBQ1IsT0FBTyxFQUFFO1FBQ1AsYUFBYSxFQUFFLElBQUk7S0FDcEI7SUFJRCxVQUFVLEVBQUU7UUFDVixRQUFRLEVBQUU7WUFDUixJQUFJLEVBQUUsTUFBTTtZQUNaLEtBQUssRUFBRSxFQUFFO1NBQ1Y7UUFDRCxLQUFLLEVBQUU7WUFDTCxJQUFJLEVBQUUsTUFBTTtZQUNaLEtBQUssRUFBRSxFQUFFO1NBQ1Y7UUFDRCxVQUFVLEVBQUU7WUFDVixJQUFJLEVBQUUsTUFBTTtZQUNaLEtBQUssRUFBRSxFQUFFO1NBQ1Y7UUFDRCxLQUFLLEVBQUU7WUFDTCxJQUFJLEVBQUUsTUFBTTtZQUNaLEtBQUssRUFBRSxFQUFFO1NBQ1Y7UUFDRCxJQUFJLEVBQUU7WUFDSixJQUFJLEVBQUUsT0FBTztZQUNiLEtBQUssRUFBRSxJQUFJO1NBQ1o7UUFDRCxPQUFPLEVBQUU7WUFDUCxJQUFJLEVBQUUsT0FBTztZQUNiLEtBQUssRUFBRSxLQUFLO1NBQ2I7UUFDRCxVQUFVLEVBQUU7WUFDVixJQUFJLEVBQUUsT0FBTztZQUNiLEtBQUssRUFBRSxLQUFLO1NBQ2I7UUFDRCxRQUFRLEVBQUU7WUFFUixJQUFJLEVBQUUsT0FBTztZQUNiLEtBQUssRUFBRSxJQUFJO1NBQ1o7UUFDRCxJQUFJLEVBQUU7WUFFSixJQUFJLEVBQUUsT0FBTztZQUNiLEtBQUssRUFBRSxJQUFJO1lBQ1gsUUFBUSxFQUFFLGFBQWE7U0FDeEI7UUFFRCxLQUFLLEVBQUU7WUFDTCxJQUFJLEVBQUUsTUFBTTtZQUNaLEtBQUssRUFBRSxDQUFDO1NBQ1Q7S0FDRjtJQUlELElBQUksRUFBRTtRQUNKLFlBQVksRUFBRSxFQUFFO0tBQ2pCO0lBQ0QsU0FBUyxFQUFFO1FBQ1QsUUFBUTtZQUNOLE1BQU0sSUFBSSxHQUFHLEVBQUUsQ0FBQywrQkFBK0IsRUFBRSxDQUFBO1lBQ2pELEVBQUUsQ0FBQyxhQUFhLENBQUM7Z0JBQ2YsT0FBTyxFQUFFLENBQUMsR0FBRyxFQUFFLEVBQUU7b0JBQ2YsTUFBTSxTQUFTLEdBQUcsR0FBRyxDQUFDLFFBQVEsS0FBSyxTQUFTLENBQUE7b0JBQzVDLE1BQU0sVUFBVSxHQUFHLEdBQUcsQ0FBQyxRQUFRLEtBQUssVUFBVSxDQUFBO29CQUM5QyxJQUFJLENBQUMsT0FBTyxDQUFDO3dCQUNYLEdBQUcsRUFBRSxDQUFDLFNBQVM7d0JBQ2YsaUJBQWlCLEVBQUUsa0JBQWtCLEdBQUcsQ0FBQyxXQUFXLEdBQUcsSUFBSSxDQUFDLElBQUksSUFBSTt3QkFDcEUsU0FBUyxFQUFFLFVBQVUsR0FBRyxDQUFDLFdBQVcsR0FBRyxJQUFJLENBQUMsSUFBSyxJQUFJO3dCQUNyRCxXQUFXLEVBQUUsVUFBVSxJQUFJLFNBQVMsQ0FBQyxDQUFDLENBQUMsZ0NBQWdDLEdBQUcsQ0FBQyxRQUFRLENBQUMsR0FBRyxxQkFBcUIsR0FBRyxDQUFDLFFBQVEsQ0FBQyxHQUFHLElBQUksQ0FBQyxDQUFDLENBQUMsRUFBRTtxQkFDdEksQ0FBQyxDQUFBO2dCQUNKLENBQUM7YUFDRixDQUFDLENBQUE7UUFDSixDQUFDO0tBQ0Y7SUFJRCxPQUFPLEVBQUU7UUFDUCxXQUFXLENBQUMsSUFBYTtZQUN2QixNQUFNLFFBQVEsR0FBRyxJQUFJLENBQUMsSUFBSSxDQUFDLFFBQVEsQ0FBQTtZQUNuQyxJQUFJLFlBQVksR0FBRyxFQUFFLENBQUE7WUFDckIsSUFBSSxRQUFRLEVBQUU7Z0JBQ1osWUFBWSxHQUFHLFlBQ2IsSUFBSSxDQUFDLENBQUMsQ0FBQyxHQUFHLENBQUMsQ0FBQyxDQUFDLEdBQ2YsMkJBQTJCLENBQUE7YUFDNUI7aUJBQU07Z0JBQ0wsWUFBWSxHQUFHLFlBQVksSUFBSSxDQUFDLENBQUMsQ0FBQyxFQUFFLENBQUMsQ0FBQyxDQUFDLE1BQU0sRUFBRSxDQUFBO2FBQ2hEO1lBQ0QsSUFBSSxDQUFDLE9BQU8sQ0FBQztnQkFDWCxZQUFZO2FBQ2IsQ0FBQyxDQUFBO1FBQ0osQ0FBQztRQUNELElBQUk7WUFDRixNQUFNLElBQUksR0FBRyxJQUFJLENBQUMsSUFBSSxDQUFBO1lBQ3RCLElBQUksSUFBSSxDQUFDLEtBQUssRUFBRTtnQkFDZCxFQUFFLENBQUMsWUFBWSxDQUFDO29CQUNkLEtBQUssRUFBRSxJQUFJLENBQUMsS0FBSztpQkFDbEIsQ0FBQyxDQUFBO2FBQ0g7WUFDRCxJQUFJLENBQUMsWUFBWSxDQUFDLE1BQU0sRUFBRSxFQUFFLEtBQUssRUFBRSxJQUFJLENBQUMsS0FBSyxFQUFFLEVBQUUsRUFBRSxDQUFDLENBQUE7UUFDdEQsQ0FBQztLQUNGO0NBQ0YsQ0FBQyxDQUFBIiwic291cmNlc0NvbnRlbnQiOlsiQ29tcG9uZW50KHtcclxuICBvcHRpb25zOiB7XHJcbiAgICBtdWx0aXBsZVNsb3RzOiB0cnVlIC8vIOWcqOe7hOS7tuWumuS5ieaXtueahOmAiemhueS4reWQr+eUqOWkmnNsb3TmlK/mjIFcclxuICB9LFxyXG4gIC8qKlxyXG4gICAqIOe7hOS7tueahOWxnuaAp+WIl+ihqFxyXG4gICAqL1xyXG4gIHByb3BlcnRpZXM6IHtcclxuICAgIGV4dENsYXNzOiB7XHJcbiAgICAgIHR5cGU6IFN0cmluZyxcclxuICAgICAgdmFsdWU6ICcnXHJcbiAgICB9LFxyXG4gICAgdGl0bGU6IHtcclxuICAgICAgdHlwZTogU3RyaW5nLFxyXG4gICAgICB2YWx1ZTogJydcclxuICAgIH0sXHJcbiAgICBiYWNrZ3JvdW5kOiB7XHJcbiAgICAgIHR5cGU6IFN0cmluZyxcclxuICAgICAgdmFsdWU6ICcnXHJcbiAgICB9LFxyXG4gICAgY29sb3I6IHtcclxuICAgICAgdHlwZTogU3RyaW5nLFxyXG4gICAgICB2YWx1ZTogJydcclxuICAgIH0sXHJcbiAgICBiYWNrOiB7XHJcbiAgICAgIHR5cGU6IEJvb2xlYW4sXHJcbiAgICAgIHZhbHVlOiB0cnVlXHJcbiAgICB9LFxyXG4gICAgbG9hZGluZzoge1xyXG4gICAgICB0eXBlOiBCb29sZWFuLFxyXG4gICAgICB2YWx1ZTogZmFsc2VcclxuICAgIH0sXHJcbiAgICBob21lQnV0dG9uOiB7XHJcbiAgICAgIHR5cGU6IEJvb2xlYW4sXHJcbiAgICAgIHZhbHVlOiBmYWxzZSxcclxuICAgIH0sXHJcbiAgICBhbmltYXRlZDoge1xyXG4gICAgICAvLyDmmL7npLrpmpDol4/nmoTml7blgJlvcGFjaXR55Yqo55S75pWI5p6cXHJcbiAgICAgIHR5cGU6IEJvb2xlYW4sXHJcbiAgICAgIHZhbHVlOiB0cnVlXHJcbiAgICB9LFxyXG4gICAgc2hvdzoge1xyXG4gICAgICAvLyDmmL7npLrpmpDol4/lr7zoiKrvvIzpmpDol4/nmoTml7blgJluYXZpZ2F0aW9uLWJhcueahOmrmOW6puWNoOS9jei/mOWcqFxyXG4gICAgICB0eXBlOiBCb29sZWFuLFxyXG4gICAgICB2YWx1ZTogdHJ1ZSxcclxuICAgICAgb2JzZXJ2ZXI6ICdfc2hvd0NoYW5nZSdcclxuICAgIH0sXHJcbiAgICAvLyBiYWNr5Li6dHJ1ZeeahOaXtuWAme+8jOi/lOWbnueahOmhtemdoua3seW6plxyXG4gICAgZGVsdGE6IHtcclxuICAgICAgdHlwZTogTnVtYmVyLFxyXG4gICAgICB2YWx1ZTogMVxyXG4gICAgfSxcclxuICB9LFxyXG4gIC8qKlxyXG4gICAqIOe7hOS7tueahOWIneWni+aVsOaNrlxyXG4gICAqL1xyXG4gIGRhdGE6IHtcclxuICAgIGRpc3BsYXlTdHlsZTogJydcclxuICB9LFxyXG4gIGxpZmV0aW1lczoge1xyXG4gICAgYXR0YWNoZWQoKSB7XHJcbiAgICAgIGNvbnN0IHJlY3QgPSB3eC5nZXRNZW51QnV0dG9uQm91bmRpbmdDbGllbnRSZWN0KClcclxuICAgICAgd3guZ2V0U3lzdGVtSW5mbyh7XHJcbiAgICAgICAgc3VjY2VzczogKHJlcykgPT4ge1xyXG4gICAgICAgICAgY29uc3QgaXNBbmRyb2lkID0gcmVzLnBsYXRmb3JtID09PSAnYW5kcm9pZCdcclxuICAgICAgICAgIGNvbnN0IGlzRGV2dG9vbHMgPSByZXMucGxhdGZvcm0gPT09ICdkZXZ0b29scydcclxuICAgICAgICAgIHRoaXMuc2V0RGF0YSh7XHJcbiAgICAgICAgICAgIGlvczogIWlzQW5kcm9pZCxcclxuICAgICAgICAgICAgaW5uZXJQYWRkaW5nUmlnaHQ6IGBwYWRkaW5nLXJpZ2h0OiAke3Jlcy53aW5kb3dXaWR0aCAtIHJlY3QubGVmdH1weGAsXHJcbiAgICAgICAgICAgIGxlZnRXaWR0aDogYHdpZHRoOiAke3Jlcy53aW5kb3dXaWR0aCAtIHJlY3QubGVmdCB9cHhgLFxyXG4gICAgICAgICAgICBzYWZlQXJlYVRvcDogaXNEZXZ0b29scyB8fCBpc0FuZHJvaWQgPyBgaGVpZ2h0OiBjYWxjKHZhcigtLWhlaWdodCkgKyAke3Jlcy5zYWZlQXJlYS50b3B9cHgpOyBwYWRkaW5nLXRvcDogJHtyZXMuc2FmZUFyZWEudG9wfXB4YCA6IGBgXHJcbiAgICAgICAgICB9KVxyXG4gICAgICAgIH1cclxuICAgICAgfSlcclxuICAgIH0sXHJcbiAgfSxcclxuICAvKipcclxuICAgKiDnu4Tku7bnmoTmlrnms5XliJfooahcclxuICAgKi9cclxuICBtZXRob2RzOiB7XHJcbiAgICBfc2hvd0NoYW5nZShzaG93OiBib29sZWFuKSB7XHJcbiAgICAgIGNvbnN0IGFuaW1hdGVkID0gdGhpcy5kYXRhLmFuaW1hdGVkXHJcbiAgICAgIGxldCBkaXNwbGF5U3R5bGUgPSAnJ1xyXG4gICAgICBpZiAoYW5pbWF0ZWQpIHtcclxuICAgICAgICBkaXNwbGF5U3R5bGUgPSBgb3BhY2l0eTogJHtcclxuICAgICAgICAgIHNob3cgPyAnMScgOiAnMCdcclxuICAgICAgICB9O3RyYW5zaXRpb246b3BhY2l0eSAwLjVzO2BcclxuICAgICAgfSBlbHNlIHtcclxuICAgICAgICBkaXNwbGF5U3R5bGUgPSBgZGlzcGxheTogJHtzaG93ID8gJycgOiAnbm9uZSd9YFxyXG4gICAgICB9XHJcbiAgICAgIHRoaXMuc2V0RGF0YSh7XHJcbiAgICAgICAgZGlzcGxheVN0eWxlXHJcbiAgICAgIH0pXHJcbiAgICB9LFxyXG4gICAgYmFjaygpIHtcclxuICAgICAgY29uc3QgZGF0YSA9IHRoaXMuZGF0YVxyXG4gICAgICBpZiAoZGF0YS5kZWx0YSkge1xyXG4gICAgICAgIHd4Lm5hdmlnYXRlQmFjayh7XHJcbiAgICAgICAgICBkZWx0YTogZGF0YS5kZWx0YVxyXG4gICAgICAgIH0pXHJcbiAgICAgIH1cclxuICAgICAgdGhpcy50cmlnZ2VyRXZlbnQoJ2JhY2snLCB7IGRlbHRhOiBkYXRhLmRlbHRhIH0sIHt9KVxyXG4gICAgfVxyXG4gIH0sXHJcbn0pXHJcbiJdfQ==