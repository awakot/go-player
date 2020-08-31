今回主に利用するオーディオのパッケージはbeepというものです。


まずはbeepの簡単な利用方法に関して説明していきます。

beepのGithubページのWikiをご覧ください
https://github.com/faiface/beep/wiki/Hello,-Beep!

まずbeepでは様々な音声ファイルをDecodeすることが可能です。そしてDecodeされたかされた返り値は

```go
    f, err := os.Open("sample.wav")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
    }

    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

    speaker.Play(streamer)
    select {}

```

ここでは、WAVファイルを利用するので```github.com/faiface/beep/wav```をimportしてください。
注目すべきはDecode関数が返す値です。

1つ目に返しているのstreamはbeepのStreamインターフェイスです。これは音声ファイルの再生などに用いられるものになります。実際にこのInterfaceの中身を見てみましょう。

そして2つ目に返しているのはformatです。これは音声ファイルの情報を持ちます。例えばその音声のサンプレートなどの情報を持ちます。


そして、次に登場するのが実際にここまで生成した要素を用いて音声を再生するspeakerパッケージです。こちらはOtoというGo言語で作成された低レベルなサウンドライブラリを用いて実際に音声を再生します。

Init関数の引数として用いられるのが先ほどDecodeして取得したformatのSamplerateとバッファサイズになります。そしてPlay関数を実行すると実際にお音声を再生することができます。


以上がbeepパッケージを用いて音を鳴らす例になります。次の章では実際にこのパッケージを用いて、コマンドライン上で利用できる簡易ミュージックプレイヤーを作成していきたいと思います。

