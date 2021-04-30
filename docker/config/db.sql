SET NAMES utf8mb4;

-- tipo_documento

SET character_set_client = utf8mb4;
CREATE TABLE IF NOT EXISTS `tipo_documento` (
  `tipo_documento_id` int NOT NULL AUTO_INCREMENT,
  `tipo_documento` varchar(100) NOT NULL,
  PRIMARY KEY (`tipo_documento_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `tipo_documento` WRITE;
UNLOCK TABLES;

-- empleos 

SET character_set_client = utf8mb4 ;
CREATE TABLE IF NOT EXISTS `empleos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `descripcion` varchar(300) NOT NULL,
  `dependencia` varchar(100) NOT NULL,
  `fecha` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

LOCK TABLES `empleos` WRITE;
UNLOCK TABLES;

-- personas_historial

SET GLOBAL time_zone ='-03:00';

SET character_set_client = utf8mb4 ;
CREATE TABLE IF NOT EXISTS `personas_historial` (
`id` int NOT NULL AUTO_INCREMENT,
`NroSerie` INT NOT NULL,
`SID` INT NOT NULL,
`SubmittedTime` DATETIME NOT NULL,
`CompletedTime` DATETIME NOT NULL,
`ModifiedTime` DATETIME NOT NULL,
`Borrador` INT NOT NULL,
`IP` VARCHAR(20) NOT NULL,
`UID` INT NOT NULL,
`NombreUsuarie` VARCHAR(100) NULL,
`Nombre` VARCHAR(100) NOT NULL,
`Apellido` VARCHAR(100) NOT NULL,
`TipoDocumento` VARCHAR(25) NULL CHECK (TipoDocumento in ('DNI', 'Sin DNI', 'Pasaporte', 'Otro')),
`NroDocumento` VARCHAR(100) NULL,
`FechaNacimiento` DATE NOT NULL,
`IdentidadGenero` VARCHAR(100) NOT NULL,
`IndicaTelefono` VARCHAR(1) NULL CHECK (IndicaTelefono in ('X')),
`IndicaEmail` VARCHAR(1) NULL CHECK (IndicaEmail in ('X')),
`CodArea` VARCHAR(6) NULL,
`Telefono` VARCHAR(20) NULL,
`IndicaOtroTelefono` VARCHAR(2) NULL CHECK (IndicaOtroTelefono in ('Si', 'No')),
`OtroCodArea` VARCHAR(6) NULL,
`OtroTelefono` VARCHAR(20) NULL,
`Email` VARCHAR(150) NULL,
`Provincia` VARCHAR(200) NOT NULL,
`Localidad` VARCHAR(200) NOT NULL,
`NivelEducativo` VARCHAR(100) NOT NULL CHECK (NivelEducativo in ('Primario', 'Secundario', 'Terciario No Universitario', 'Universitario', 'Posgrado', 'Sin estudios')),
`CompletoNivelEducativo` VARCHAR(100) NULL CHECK (CompletoNivelEducativo in ('No, se encuentra en curso', 'No, no lo completó', 'Si', 'Sin datos')),
`Carrera` VARCHAR(300) NULL,
`IndicaBeca` VARCHAR(2) NULL CHECK (IndicaBeca in ('Si', 'No')),
`InteresBeca` VARCHAR(2) NULL CHECK (InteresBeca in ('Si', 'No')),
`NombreBeca` VARCHAR(300) NULL,
`InstitucionBeca` VARCHAR(300) NULL,
`IndicaPFE` VARCHAR(2) NULL CHECK (IndicaPFE in ('Si', 'No')),
`NombrePFE` VARCHAR(300) NULL,
`LugarPFE` VARCHAR(300) NULL,
`InteresFE` VARCHAR(2) NULL CHECK (InteresFE in ('Si', 'No')),
`PaisNacimiento` VARCHAR(20) NOT NULL CHECK (PaisNacimiento in ('Argentina', 'Otro país')),
`ProvinciaNacimiento` VARCHAR(300) NULL,
`OtroPaisNacimiento` VARCHAR(300) NULL,
`InteresDPTM` VARCHAR(2) NULL CHECK (InteresDPTM in ('Si', 'No')),
`Trabaja` VARCHAR(2) NOT NULL CHECK (Trabaja in ('Si', 'No')),
`Ocupacion` VARCHAR(300) NULL,
`DescuentoJubilacion` VARCHAR(2) NULL CHECK (DescuentoJubilacion in ('Si', 'No')),
`BuscaTrabajo` VARCHAR(2) NOT NULL CHECK (BuscaTrabajo in ('Si', 'No')),
`IndicaPCL` VARCHAR(2) NULL CHECK (IndicaPCL in ('Si', 'No')),
`InteresPCL` VARCHAR(2) NULL CHECK (InteresPCL in ('Si', 'No')),
`RecibePrestacion` VARCHAR(2) NULL CHECK (RecibePrestacion in ('Si', 'No')),
`AUHijeEmbarazo` VARCHAR(1) NULL CHECK (AUHijeEmbarazo in ('X')),
`AHijeConDiscapacidad` VARCHAR(1) NULL CHECK (AHijeConDiscapacidad in ('X')),
`IFE` VARCHAR(1) NULL CHECK (IFE in ('X')),
`PPT` VARCHAR(1) NULL CHECK (PPT in ('X')),
`PrestacionPorDesempleo` VARCHAR(1) NULL CHECK (PrestacionPorDesempleo in ('X')),
`OtrosNacion` VARCHAR(1) NULL CHECK (OtrosNacion in ('X')),
`OtrosCABA` VARCHAR(1) NULL CHECK (OtrosCABA in ('X')),
`OtrosProvinciales` VARCHAR(1) NULL CHECK (OtrosProvinciales in ('X')),
`OtrosMunicipales` VARCHAR(1) NULL CHECK (OtrosMunicipales in ('X')),
`OtraAsignacion` VARCHAR(1) NULL CHECK (OtraAsignacion in ('X')),
`NombreOtraAsignacion` VARCHAR(300) NULL,
`ObraSocial` VARCHAR(2) NOT NULL CHECK (ObraSocial in ('Si', 'No')),
`HospitalPublico` VARCHAR(2) NOT NULL CHECK (HospitalPublico in ('Si', 'No')),
`CentroDeSalud` VARCHAR(2) NOT NULL CHECK (CentroDeSalud in ('Si', 'No')),
`ConsultorioAmigable` VARCHAR(2) NOT NULL CHECK (ConsultorioAmigable in ('Si', 'No')),
`ConsultorioPrivado` VARCHAR(2) NOT NULL CHECK (ConsultorioPrivado in ('Si', 'No')),
`OtroCentroDeSalud` VARCHAR(2) NOT NULL CHECK (OtroCentroDeSalud in ('Si', 'No')),
`NombreOtroCentroDeSalud` VARCHAR(300) NULL,
`IndicaAsesoramientoADSI` VARCHAR(2) NULL CHECK (IndicaAsesoramientoADSI in ('Si', 'No')),
`TratamientoHormonal` VARCHAR(1) NULL CHECK (TratamientoHormonal in ('X')),
`IntervencionCirugia` VARCHAR(1) NULL CHECK (IntervencionCirugia in ('X')),
`MedicacionEspecifica` VARCHAR(1) NULL CHECK (MedicacionEspecifica in ('X')),
`OtroAsesoramiento` VARCHAR(1) NULL CHECK (OtroAsesoramiento in ('X')),
`NombreOtroAsesoramiento` VARCHAR(300) NULL,
`InteresRCLTAP` VARCHAR(2) NULL CHECK (InteresRCLTAP in ('Si', 'No')),
`AreaPotencialYExperiencia` VARCHAR(300) NULL,
`NombreOtraArea` VARCHAR(300) NULL,
`OtraAreaPotencialYExperiencia` VARCHAR(300) NULL,
`Opcion1` VARCHAR(300) NULL,
`Opcion2` VARCHAR(300) NULL,
`Opcion3` VARCHAR(300) NULL,
`Opcion4` VARCHAR(300) NULL,
`Opcion5` VARCHAR(300) NULL,
`OtrosSaberes` VARCHAR(300) NULL,
`NombreCV` VARCHAR(300) NULL,
`TamanoArchivoCV` VARCHAR(300) NULL,
`NombreXLSX` VARCHAR(300) NOT NULL,
`LineaEnXLSX` INT NOT NULL,
`Timestamp` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;