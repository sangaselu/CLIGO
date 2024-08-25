CREATE   PROCEDURE [dbo].[NS_TMMOTIVOS_EXTORNO_PEDIDO_F](
	@id int
)
AS
	DECLARE @ErrorMessage NVARCHAR(max),
			@ErrorSeverity INT,
			@ErrorState INT;

BEGIN TRY
	IF NOT EXISTS (SELECT id FROM TMMOTIVOS_EXTORNO_PEDIDO WITH(NOLOCK) WHERE id = @id)
			THROW 51003, 'No existe ninguna registro con este identificador', 15;
	
	SELECT *
	FROM TMMOTIVOS_EXTORNO_PEDIDO t
	WHERE t.id=@id

END TRY
BEGIN CATCH
	SELECT
		@ErrorMessage='Ocurrio un Error: '+ERROR_MESSAGE() + ' LN. ' + CONVERT(NVARCHAR(255), ERROR_LINE() ) + '.',
		@ErrorSeverity=ERROR_SEVERITY(),
		@ErrorState=ERROR_STATE();
	THROW 51003, @ErrorMessage, 10
END CATCH