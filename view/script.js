$(function () {
  var ws = new WebSocket("ws://localhost:8080/ws");
  ws.onmessage = function (evt) {
    var received_msg = evt.data;

    const message = createMessage(received_msg, 'received');
    messageList.append(message);
    $.each($('.content-effect'), function (index, element) {
      if ($(element).text() == '') {
        let text = $(element).attr('data-val')// 欲顯示的文字
        typeEffect($(element), text, 80); // 呼叫函數開始顯示文字
      }


      scrollToBottom();
    });
  };

  const messageList = $('#message-list');
  const messageInput = $('#message-input');
  var intervalId
  var effectElement
  // 將文字逐個字地添加到目標元素中
  function typeEffect(element, text, delay) {
    $('#refresh').css('display', 'none')
    $('#stop').css('display', 'block')
    var i = 0;
    effectElement = element
    intervalId = setInterval(function () {
      element.append(text.charAt(i));
      i++;
      if (i >= text.length) {
        clearInterval(intervalId);
        element.removeClass('content-effect')
        $('#stop').css('display', 'none')
        $('#refresh').css('display', 'block')
      }
    }, delay);

  }
  //
  $('#stop').click(function () {
    clearInterval(intervalId);
    effectElement.removeClass('content-effect')
    $('#stop').css('display', 'none')
    $('#refresh').css('display', 'block')
  });

  $('#send-button').click(function () {
    const content = messageInput.val().trim();
    const contentView = messageInput.val().replace(/\n/g, '<br>')
    ws.send(content);
    renderMessage(content, contentView)
    $('#refresh').attr('data-content', content).attr('data-contentView', contentView)
  });

  $('#refresh').click(function () {
    const content = $(this).attr('data-content');
    const contentView = $(this).attr('data-contentView');
    renderMessage(content, contentView)
  });
  //
  function renderMessage(content, contentView) {
    if (content.length > 0) {
      const message = createMessage(contentView, 'sent');
      messageList.append(message);
      scrollToBottom();
      /*
      setTimeout(() => {//假裝延遲
        sendMessage(content).then((response) => {
          const message = createMessage(response, 'received');
          messageList.append(message);
          $.each($('.content-effect'), function (index, element) { 
              if($(element).text()==''){
                let text = $(element).attr('data-val')// 欲顯示的文字
                typeEffect($(element),text,200); // 呼叫函數開始顯示文字
              }
          });
          
          scrollToBottom();
        });
      }, "1500");
*/
    }
    messageInput.val('');
    $('.message-intro').remove()
  }
  $('#message-input').on('keydown', function (e) {
    console.log(e);
    if (e.keyCode == 13 && e.shiftKey) {
    } else if (event.keyCode === 13) {
      event.preventDefault();
      $('#send-button').trigger('click');
    }
  });

  function createMessage(content, type) {
    if (type === 'sent') {
      message = `
      <div class="message message-user">
        <div class="message-box">
          <span class="message-sender">You</span>
          <span class="message-content">${content}</span>
        </div>
      </div>`;
    } else {
      message = `
      <div class="message message-owner">
        <div class="message-box">
          <span class="message-sender">AI</span>
          <span class="message-content content-effect" data-val="${content}"></span>
        </div>
      </div>`;
    }

    return message;
  }

  async function sendMessage(messageText) {
    /*const response = await $.ajax({
      url: 'https://api.openai.com/v1/engines/davinci-codex/completions',
      type: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + 'INSERT_YOUR_API_KEY_HERE',
      },
      data: JSON.stringify({
        prompt: 'AI: ' + messageText,
        max_tokens: 100,
      }),
    });

    return response.choices[0].text.trim();*/
    return '以下是使用jQuery生成隨機5位數字的範例程序代碼'
  }

  function scrollToBottom() {
    messageList.scrollTop(messageList.prop('scrollHeight'));
  }
  $('#openMenu').click(function () {
    $('body').addClass('activeMenu')
  });
  $(document).click(function (event) {
    if ($("body").hasClass("activeMenu")) {
      if (!($(event.target).closest('#openMenu').length || $(event.target).closest('#menu').length)) {
        // 執行你想要的程式碼
        $('body').removeClass('activeMenu')
      }
    }
  });
});