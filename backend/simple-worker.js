// Cloudflare Workers用のシンプルなHello World API
export default {
  async fetch(request, env, ctx) {
    // CORS設定
    const corsHeaders = {
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Methods': 'GET, POST, PUT, DELETE, OPTIONS',
      'Access-Control-Allow-Headers': 'Content-Type, Authorization',
    };

    // OPTIONSリクエストの処理
    if (request.method === 'OPTIONS') {
      return new Response(null, {
        status: 204,
        headers: corsHeaders,
      });
    }

    const url = new URL(request.url);
    
    // パスに基づいて異なるレスポンスを返す
    switch (url.pathname) {
      case '/':
        return new Response(
          JSON.stringify({
            message: 'Hello World from FlowGrid API!',
            status: 'success',
            version: '1.0.0',
          }),
          {
            status: 200,
            headers: {
              'Content-Type': 'application/json',
              ...corsHeaders,
            },
          }
        );
      case '/health':
        return new Response(
          JSON.stringify({
            status: 'healthy',
          }),
          {
            status: 200,
            headers: {
              'Content-Type': 'application/json',
              ...corsHeaders,
            },
          }
        );
      default:
        return new Response(
          JSON.stringify({
            error: 'Not Found',
          }),
          {
            status: 404,
            headers: {
              'Content-Type': 'application/json',
              ...corsHeaders,
            },
          }
        );
    }
  },
};
