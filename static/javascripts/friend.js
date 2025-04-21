$(document).ready(function() {
    // 点击推荐好友按钮
    $(document).on('click', '.a-recommend-list', function(e) {
        e.stopPropagation();
        var popover = $(this).closest('.recommendlist').find('.popover');
        $('.popover').not(popover).hide();
        popover.toggle();
        if (popover.is(':visible')) {
            loadRecommendFriends();
        }
    });

    // 加载推荐好友列表
    function loadRecommendFriends() {
        $.ajax({
            url: '/friend/recommend',
            type: 'POST',
            success: function(res) {
                if (res.code === 0) {
                    var recommendList = res.data;
                    var html = '';
                    recommendList.forEach(function(user) {
                        html += '<li class="li-friend-item" data-uid="' + user.user_id + '" data-username="' + user.username + '" data-avatar_id="' + user.avatar_id + '">';
                        html += '<img src="/static/images/user/' + user.avatar_id + '.png" alt="portrait_' + user.avatar_id + '">';
                        html += '<b>' + user.username + '</b>';
                        html += '<button class="btn btn-primary btn-sm add-friend-btn" style="float:right;margin-top:10px;">添加好友</button>';
                        html += '</li>';
                    });
                    $('.ul-recommend-list').html(html || '<li>暂无推荐好友</li>');
                } else {
                    layer.msg(res.msg);
                }
            },
            error: function() {
                layer.msg('获取推荐好友列表失败，请重试');
            }
        });
    }

    // 点击头像显示添加好友按钮
    $(document).on('click', '.user-avatar', function(e) {
        e.stopPropagation();
        var popup = $(this).siblings('.add-friend-popup');
        $('.add-friend-popup').not(popup).hide();
        popup.toggle();
    });

    // 点击其他地方隐藏添加好友按钮
    $(document).on('click', function() {
        $('.add-friend-popup').hide();
    });

    // 点击添加好友按钮
    $(document).on('click', '.add-friend-btn', function(e) {
        e.stopPropagation();
        var userAvatar = $(this).closest('li').find('.user-avatar');
        var toUserId = userAvatar.data('uid');

        $.ajax({
            url: '/friend/add',
            type: 'POST',
            data: {
                to_user_id: toUserId
            },
            success: function(res) {
                if (res.code === 0) {
                    layer.msg('添加好友成功');
                    loadFriendList(); // 刷新好友列表
                } else {
                    layer.msg(res.msg);
                }
            },
            error: function() {
                layer.msg('添加好友失败，请重试');
            }
        });

        $(this).closest('.add-friend-popup').hide();
    });

    // 点击好友列表按钮
    $(document).on('click', '.a-friend-list', function(e) {
        e.stopPropagation();
        var popover = $(this).closest('.friendlist').find('.popover');
        $('.popover').not(popover).hide();
        popover.toggle();
        if (popover.is(':visible')) {
            loadFriendList();
        }
    });

    // 点击其他地方隐藏好友列表
    $(document).on('click', function() {
        $('.friendlist .popover').hide();
    });

    // 加载好友列表
    function loadFriendList() {
        $.ajax({
            url: '/friend/list',
            type: 'GET',
            success: function(res) {
                if (res.code === 0) {
                    var friendList = res.data;
                    var html = '';
                    friendList.forEach(function(friend) {
                        html += '<li class="li-friend-item" data-uid="' + friend.user_id + '" data-username="' + friend.username + '" data-avatar_id="' + friend.avatar_id + '">';
                        html += '<img src="/static/images/user/' + friend.avatar_id + '.png" alt="portrait_' + friend.avatar_id + '">';
                        html += '<b>' + friend.username + '</b>';
                        html += '</li>';
                    });
                    $('.ul-friend-list').html(html || '<li>暂无好友</li>');
                } else {
                    layer.msg(res.msg);
                }
            },
            error: function() {
                layer.msg('获取好友列表失败，请重试');
            }
        });
    }
});