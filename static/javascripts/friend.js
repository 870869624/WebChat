// 好友相关的交互逻辑
$(document).ready(function() {
    // 点击用户头像显示添加好友按钮
    $(document).on('click', '.user-avatar', function(e) {
        e.stopPropagation(); // 阻止事件冒泡
        $('.add-friend-popup').hide(); // 隐藏其他弹出层
        $(this).siblings('.add-friend-popup').show();
    });

    // 点击其他地方隐藏添加好友按钮
    $(document).on('click', function() {
        $('.add-friend-popup').hide();
    });

    // 点击添加好友按钮
    $(document).on('click', '.add-friend-btn', function(e) {
        e.stopPropagation();
        const toUserId = $(this).closest('li').find('.user-avatar').data('uid');
        const toUsername = $(this).closest('li').find('.user-avatar').data('username');
        
        // 发送添加好友请求
        $.ajax({
            url: '/friend/add',
            type: 'POST',
            data: {
                to_user_id: toUserId
            },
            success: function(res) {
                if (res.code === 0) {
                    layer.msg('已发送好友请求给 ' + toUsername);
                } else {
                    layer.msg(res.msg || '发送好友请求失败');
                }
            },
            error: function() {
                layer.msg('网络错误，请稍后重试');
            }
        });

        // 隐藏弹出层
        $('.add-friend-popup').hide();
    });
});