Imports System.IO
Imports System.Net
Imports System.Windows.Forms.VisualStyles.VisualStyleElement.Tab
Imports Newtonsoft.Json
Imports Newtonsoft.Json.Linq
Public Class Form1

    Private Sub Form1_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        ' Menginisialisasi DataGrid
        Dim grid As New DataGridView()
        grid.Dock = DockStyle.Fill
        Me.Controls.Add(grid)

        ' Mengirim permintaan HTTP ke API dan mendapatkan respons JSON
        Dim apiUrl As String = "http://localhost:8080/api/nasabah"
        Dim request As WebRequest = WebRequest.Create(apiUrl)
        Dim response As WebResponse = request.GetResponse()
        Dim dataStream As Stream = response.GetResponseStream()
        Dim reader As New StreamReader(dataStream)
        Dim jsonResponse As String = reader.ReadToEnd()

        ' Parsing JSON menjadi daftar objek DataModel
        Dim data As List(Of DataModel) = JsonConvert.DeserializeObject(Of List(Of DataModel))(jsonResponse)

        ' Mengisi DataGrid dengan data yang sudah diparse
        grid.DataSource = data

        ' Membersihkan sumber daya
        reader.Close()
        dataStream.Close()
        response.Close()
    End Sub
    Private Sub Button1_Click(sender As Object, e As EventArgs) Handles Button1.Click
        Dim Client As New WebClient
        Dim Result = Client.DownloadString("http://localhost:8080/api/nasabah")
        Dim Json = JToken.Parse(Result)
        MsgBox(Json(0)("alamat_nasabah"))
    End Sub

End Class
